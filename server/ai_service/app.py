from fastapi import FastAPI
from pydantic import BaseModel
import cv2
import easyocr
import numpy as np
import re
from difflib import SequenceMatcher
from pdf2image import convert_from_path
from PIL import Image
import torch
import imagehash
from transformers import CLIPProcessor, CLIPModel
import os

app = FastAPI()

device = "cuda" if torch.cuda.is_available() else "cpu"

# Lightweight CLIP model
model_name = "openai/clip-vit-base-patch32"

processor = CLIPProcessor.from_pretrained(model_name)
model = CLIPModel.from_pretrained(model_name).to(device)

# Initialize ocr 
ocr_reader = easyocr.Reader(['en'])

# Known OCR confusion pairs
OCR_CORRECTIONS = {
    '0': 'O', 'O': '0',
    '1': 'I', 'I': '1',
    '1': 'Z', 'Z': '1',   # your exact case
    '8': 'B', 'B': '8',
    '5': 'S', 'S': '5',
    '2': 'Z', 'Z': '2',
}

GST_PATTERN = r"\d{2}[A-Z]{5}\d{4}[A-Z]\d[A-Z0-9]{2}"  # relaxed last 2 chars

class AnalyzeRequest(BaseModel):
    image_path: str
    cause_text: str
    previous_phashes: list[str] = []

class ReceiptRequest(BaseModel):
    receipt_path: str
    claimed_amount: float

# Helper functions for gst validation 
def correct_ocr_gst(raw_gst):
    """Try to fix common OCR mistakes in GST numbers."""
    corrected = list(raw_gst.upper())
    print (f"Corrected gst: {corrected}")
    
    # Position 12 should always be 'Z' in valid GST
    if len(corrected) >= 14 and corrected[13] != 'Z':
        if corrected[13] in ('1', '2', '7'):  # common Z misreads
            corrected[13] = 'Z'
    
    return "".join(corrected)

def extract_and_correct_gst(text):
    # Relaxed pattern - last chars can be anything alphanumeric
    raw_match = re.search(r"\d{2}[A-Z0-9]{5}\d{4}[A-Z0-9]\d[A-Z0-9]{2}", text)
    if not raw_match:
        return None, None
    
    raw_gst = raw_match.group()
    corrected_gst = correct_ocr_gst(raw_gst)
    
    return raw_gst, corrected_gst

def validate_gst(gst):
    if gst is None:
        return False
    # Strict format after correction
    pattern = r"\d{2}[A-Z]{5}\d{4}[A-Z]\d[Z][A-Z0-9]"
    return bool(re.match(pattern, gst))

# <-- Helper Functions For analyze-Reciept -->
def validate_file(path):

    if not os.path.exists(path):
        return False, "file_not_found"

    allowed = ["jpg", "jpeg", "png", "pdf"]

    ext = path.split(".")[-1].lower()
    if ext not in allowed:
        return False, "invalid_format"

    size = os.path.getsize(path)
    if size > 5 * 1024 * 1024:
        return False, "file_too_large"

    return True, None

def convert_pdf(path):

    if path.endswith(".pdf"):
        pages = convert_from_path(path)
        temp_img = "temp_receipt.jpg"
        pages[0].save(temp_img, "JPEG")
        return temp_img

    return path

def preprocess_receipt(path):

    img = cv2.imread(path)
    gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
    blur = cv2.GaussianBlur(gray, (5,5), 0)
    # sharpen
    # kernel = np.array([[0,-1,0],
    #                    [-1,5,-1],
    #                    [0,-1,0]])

    # sharp = cv2.filter2D(gray, -1, kernel)
    thresh = cv2.adaptiveThreshold(
        blur,
        # sharp,
        255,
        cv2.ADAPTIVE_THRESH_GAUSSIAN_C,
        cv2.THRESH_BINARY,
        11,
        2
    )

    return thresh

def extract_receipt_text(img):

    results = ocr_reader.readtext(img)
    text = ""
    for r in results:
        text += r[1] + " "

    return text

def normalize_ocr_text(text):

    text = text.upper()

    # remove currency symbols
    text = text.replace("₹", "")
    text = text.replace("RS.", "")
    text = text.replace("RS", "")

    # replace OCR punctuation mistakes
    text = text.replace(";", ".")
    # text = text.replace(",", ".")
    text = text.replace(":", " ")

    # remove duplicated dots
    text = re.sub(r"\.{2,}", ".", text)

    return text


def clean_amount(value):

    value = value.strip()

    # case 1: 6,440.00 → remove thousand comma
    if "," in value and "." in value:
        value = value.replace(",", "")

    # case 2: 23779,00 → convert comma decimal
    elif "," in value:
        value = value.replace(",", ".")

    # remove OCR junk
    value = re.sub(r"[^\d.]", "", value)
    
    return float(value)

def parse_receipt(text, claimed_amount=None):

    # gst = None
    amount = None

    # -------- GST Detection --------
    # gst_match = re.search(r"\d{2}[A-Z]{4,5}\d{4}[A-Z0-9]{3,5}", text)

    # if gst_match:
    #     gst = gst_match.group()
    raw_gst, corrected_gst = extract_and_correct_gst(text)
    
    gst_was_corrected = (raw_gst != corrected_gst)

    # -------- Amount Keywords Priority --------
    keywords = [
        "GRAND TOTAL",
        "TOTAL AMOUNT",
        "NET AMOUNT",
        "AMOUNT PAID",
        "PAYABLE",
        "TOTAL"
    ]

    candidates = []

    for key in keywords:

        # Replace spaces in keyword with \s+ to match one or more spaces
        key_pattern = r"\s+".join(re.escape(word) for word in key.split())
        pattern = rf"\b{key_pattern}\b[^\d]{{0,20}}(\d{{1,3}}(?:[.,]\d{{3}})*(?:[.,]\d{{2}}))['\"\-:;= ]{{0,5}}"

        match = re.search(pattern, text)

        if match:
            raw_value = match.group(1)
            print(f"Raw value detected: {raw_value}")

            try:
                value = clean_amount(raw_value)
                print(f"Cleaned value: {value}")

                # If leading digit is '2' and stripping it gives a value close to claimed_amount, use corrected value
                if claimed_amount:
                    str_int = str(int(value))
                    if str_int.startswith("2") and len(str_int) > 1:
                        decimal_part = round((value % 1) * 100)
                        corrected = float(f"{str_int[1:]}.{decimal_part:02d}")
                        if abs(corrected - claimed_amount) < 20:
                            print(f"[{key}] OCR ₹->2 fix applied: {value} -> {corrected}")
                            candidates.append((key, corrected))
                            continue
                
                # reject unrealistic values (₹ OCR error)
                if value > claimed_amount * 5:
                    continue                
                candidates.append((key, value))
            except Exception as e:
                print(f"[{key}] Error parsing value: {e}")
 
    print(f"Candidates: {candidates}")

    # -------- Check candidates against claimed amount --------
    if claimed_amount is not None:

        for key, value in candidates:

            if abs(value - claimed_amount) < 20:
                return value, corrected_gst, gst_was_corrected
            
            # fix OCR ₹ -> 2 error
            # if str(int(value)).startswith("2"):
            #     corrected = float(str(int(value))[1:])
            #     if abs(corrected - claimed_amount) < 20:
            #         return corrected, gst

    # -------- If no match found choose first priority keyword --------
    if candidates:
        return candidates[0][1], corrected_gst, gst_was_corrected

    # -------- Fallback: largest detected number --------
    numbers = re.findall(r"\d+[.,]\d{2}", text)

    values = []

    for n in numbers:
        try:
            values.append(clean_amount(n))
        except:
            pass

    if values:
        if claimed_amount:
            corrected_values = []
            for v in values:
                str_int = str(int(v))
                if str_int.startswith("2") and len(str_int) > 1:
                    decimal_part = round((v % 1) * 100)
                    corrected = float(f"{str_int[1:]}.{decimal_part:02d}")
                    if abs(corrected - claimed_amount) < 20:
                        corrected_values.append(corrected)
                        continue
                corrected_values.append(v)
            values = corrected_values
        amount = max(values)
        return amount, corrected_gst, gst_was_corrected

    return None, corrected_gst, gst_was_corrected

# def validate_gst(gst):

#     if gst is None:
#         return False
#     pattern = r"\d{2}[A-Z]{5}\d{4}[A-Z]\d[Z][A-Z0-9]"
#     return bool(re.match(pattern, gst))

def tamper_detection(path):

    original = cv2.imread(path)
    temp = "compressed.jpg"
    cv2.imwrite(temp, original, [cv2.IMWRITE_JPEG_QUALITY, 90])
    compressed = cv2.imread(temp)
    diff = cv2.absdiff(original, compressed)
    score = np.mean(diff)
    # print(f"Tamper score:{score}")
    return score


@app.get("/health")
def health():
    return {"status": "ok"}


@app.post("/analyze")
def analyze(data: AnalyzeRequest):

    if not os.path.exists(data.image_path):
        return {"error": "Image not found"}

    flags = []
    score = 0

    # -------- OpenCV Blur Detection --------
    img = cv2.imread(data.image_path)
    gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)

    blur_score = cv2.Laplacian(gray, cv2.CV_64F).var()
    print(f"Blur score: {blur_score}")

    if blur_score < 50:
        flags.append("blurry_image")
    else:
        score += 20

    # -------- Perceptual Hash Duplicate Detection --------
    pil_img = Image.open(data.image_path).convert("RGB")

    current_phash = imagehash.phash(pil_img)

    for prev in data.previous_phashes:
        prev_hash = imagehash.hex_to_hash(prev)

        # Hamming distance
        distance = current_phash - prev_hash

        if distance < 5:
            flags.append("Duplicate image")
            break

    # After verified images, store the hash in your database.Store: Cause ID, Image URL, Perceptual Hash, CLIP Embedding
    data.previous_phashes.append(str(current_phash))
    print(f"Previous phashes: {data.previous_phashes}")

    # -------- CLIP Semantic Similarity --------
    image = Image.open(data.image_path)

    inputs = processor(
        text=[data.cause_text],
        images=image,
        return_tensors="pt",
        padding=True
    ).to(device)

    outputs = model(**inputs)

    image_embeds = outputs.image_embeds
    text_embeds = outputs.text_embeds

    # cosine similarity
    similarity = torch.cosine_similarity(image_embeds, text_embeds).item()

    semantic_score = int((similarity + 1) / 2 * 80)  # normalize -1..1 → 0..80
    print(f"Semantic similarity: {similarity:.4f}, score: {semantic_score}")
    score += semantic_score


    # -------- Final Decision --------
    if score >= 70:
        status = "verified"
    elif score >= 40:
        status = "review"
    else:
        status = "rejected"

    return {
        "final_score": score,
        "validation_status": status,
        "flags": flags
    }

# <-- analyze-reciept Endpoint -->
@app.post("/analyze-receipt")
def analyze_receipt(data: ReceiptRequest):

    valid, error = validate_file(data.receipt_path)

    if not valid:
        return {"error": error}

    flags = []
    score = 0

    # convert pdf if needed
    path = convert_pdf(data.receipt_path)

    # preprocessing
    processed = preprocess_receipt(path)

    # OCR
    text = extract_receipt_text(processed)
    print(f"Extracted OCR Text:{text}")

    # Normalize text
    text = normalize_ocr_text(text)
    print(f"Normalize OCR Text:{text}")

    # parse fields
    amount, gst, gst_corrected = parse_receipt(text, data.claimed_amount)
    print(f"Amount:{amount}, GST: {gst}")

    # GST validation
    gst_valid = validate_gst(gst)

    if gst_valid:
        score += 25
        if gst_corrected:
            flags.append("gst_ocr_corrected")  # soft flag, not penalized
    else:
        flags.append("invalid_gst")

    # tamper detection
    tamper_score = tamper_detection(path)

    if tamper_score < 10:
        score += 25
    else:
        flags.append("possible_tampering")

    # expense matching
    if amount is not None:

        if abs(amount - data.claimed_amount) < 10:
            score += 25
        else:
            flags.append("amount_mismatch")

    else:
        flags.append("amount_not_detected")

    # vendor reputation placeholder
    if gst_valid:
        score += 25

    # final status
    print(f"Score: {score}")
    if score >= 70:
        status = "verified"
    elif score >= 40:
        status = "review"
    else:
        status = "rejected"

    return {
        "detected_amount": amount,
        "gst_number": gst,
        "gst_ocr_corrected": gst_corrected,  #  expose to frontend  
        "tamper_score": tamper_score,
        "receipt_score": score,
        "status": status,
        "flags": flags
    }