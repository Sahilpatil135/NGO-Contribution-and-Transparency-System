from fastapi import FastAPI
from pydantic import BaseModel
import cv2
import easyocr
import numpy as np
import re
from difflib import SequenceMatcher
from pdf2image import convert_from_path
from PIL import Image
# from ultralytics import YOLO
import torch
import imagehash
from transformers import CLIPProcessor, CLIPModel
import os

app = FastAPI()

device = "cuda" if torch.cuda.is_available() else "cpu"

# yolo_model = YOLO("yolov8n.pt")

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
    # products: list[str] = []
    previous_phashes: list[str] = []

class ReceiptRequest(BaseModel):
    receipt_path: str    
    claimed_amount: float

# ─────────────────────────────────────────────
#  Helper functions for GST validation
# ─────────────────────────────────────────────
def preprocess_gst_text(text):

    text = text.upper()

    # Fix common OCR mistakes
    text = text.replace("O", "0")  # O → 0
    text = text.replace("I", "1")  # I → 1

    return text

def correct_ocr_gst(raw_gst):
    # Try to fix common OCR mistakes in GST numbers.
    corrected = list(raw_gst.upper())
    print (f"Corrected gst: {corrected}")
    
    # Position 14 should always be 'Z' in valid GST
    if len(corrected) >= 14 and corrected[13] != 'Z':
        if corrected[13] in ('1', '2', '7'):  # common Z misreads
            corrected[13] = 'Z'
    
    return "".join(corrected)

def extract_and_correct_gst(text):
    # Relaxed pattern - last chars can be anything alphanumeric
    text = preprocess_gst_text(text)
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

# ─────────────────────────────────────────────
#  Helper functions for /analyze-receipt
# ─────────────────────────────────────────────
def ensure_jpg(path):
    ext = path.split(".")[-1].lower()

    if ext in ["jfif", "jpeg", "png"]:
        try:
            img = Image.open(path).convert("RGB")
            new_path = path.rsplit(".", 1)[0] + ".jpg"
            img.save(new_path, "JPEG", quality=90)
            return new_path
        except Exception as e:
            print(f"Conversion error: {e}")
            return path

    return path

def validate_file(path):

    if not os.path.exists(path):
        return False, "file_not_found"

    allowed = ["jpg", "jpeg", "png", "pdf", "jfif"]

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
    thresh = cv2.adaptiveThreshold(
        blur,        
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
        text += r[1] + "\n"

    return text

def normalize_ocr_text(text):

    text = text.upper()

    # remove currency symbols
    text = text.replace("₹", "")
    text = text.replace("RS.", "")   

    # replace OCR punctuation mistakes
    text = text.replace(";", ".")    

    # remove duplicated dots
    text = re.sub(r"\.{2,}", ".", text)

    return text


def clean_amount(value):

    value = value.strip()

    # keep only digits, comma and dot
    value = re.sub(r"[^\d.,]", "", value)

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
    
    # -------- GST Detection --------
    raw_gst, corrected_gst = extract_and_correct_gst(text)
    
    gst_was_corrected = (raw_gst != corrected_gst)

    # -------- Amount Keywords Priority --------
    # keywords = [
    #     "GRAND TOTAL",
    #     "TOTAL AMOUNT",
    #     "NET AMOUNT",
    #     "AMOUNT PAID",
    #     "PAYABLE",
    #     "TOTAL"
    # ]

    # candidates = []

    # for key in keywords:

    #     # Replace spaces in keyword with \s+ to match one or more spaces
    #     key_pattern = r"\s+".join(re.escape(word) for word in key.split())
    #     pattern = rf"\b{key_pattern}\b[^\d]{{0,20}}(\d{{1,3}}(?:[.,]\d{{3}})*(?:[.,]\d{{2}}))['\"\-:;= ]{{0,5}}"

    #     match = re.search(pattern, text)

    #     if match:
    #         raw_value = match.group(1)
    #         print(f"Raw value detected: {raw_value}")

    #         try:
    #             value = clean_amount(raw_value)
    #             print(f"Cleaned value: {value}")

    #             # If leading digit is '2' and stripping it gives a value close to claimed_amount, use corrected value
    #             if claimed_amount:
    #                 str_int = str(int(value))
    #                 if str_int.startswith("2") and len(str_int) > 1:
    #                     decimal_part = round((value % 1) * 100)
    #                     corrected = float(f"{str_int[1:]}.{decimal_part:02d}")
    #                     if abs(corrected - claimed_amount) < 20:
    #                         print(f"[{key}] OCR ₹->2 fix applied: {value} -> {corrected}")
    #                         candidates.append((key, corrected))
    #                         continue
                
    #             # reject unrealistic values (₹ OCR error)
    #             if value > claimed_amount * 5:
    #                 continue                
    #             candidates.append((key, value))
    #         except Exception as e:
    #             print(f"[{key}] Error parsing value: {e}")
 
    # print(f"Candidates: {candidates}")

    # # -------- Check candidates against claimed amount --------
    # if claimed_amount is not None:

    #     for key, value in candidates:

    #         if abs(value - claimed_amount) < 20:
    #             return value, corrected_gst, gst_was_corrected            

    # # -------- If no match found choose first priority keyword --------
    # if candidates:
    #     return candidates[0][1], corrected_gst, gst_was_corrected

    # # -------- Fallback: largest detected number --------
    # numbers = re.findall(r"\d+[.,]\d{2}", text)

    # values = []

    # for n in numbers:
    #     try:
    #         values.append(clean_amount(n))
    #     except:
    #         pass

    # if values:
    #     if claimed_amount:
    #         corrected_values = []
    #         for v in values:
    #             str_int = str(int(v))
    #             if str_int.startswith("2") and len(str_int) > 1:
    #                 decimal_part = round((v % 1) * 100)
    #                 corrected = float(f"{str_int[1:]}.{decimal_part:02d}")
    #                 if abs(corrected - claimed_amount) < 20:
    #                     corrected_values.append(corrected)
    #                     continue
    #             corrected_values.append(v)
    #         values = corrected_values
    #     amount = max(values)
    #     return amount, corrected_gst, gst_was_corrected

    # regex 1 → thousand formatted numbers
    matches1 = re.findall(r"\d+[.,]\d{2}", text)
    # regex 2 → general decimal numbers
    matches2 = re.findall(r"\d{1,3}(?:,\d{3})*(?:\.\d{2})", text)

    matches = list(set(matches1 + matches2))  # combine & remove duplicates

    amounts = []

    for m in matches:
        try:
            # val = float(m.replace(",", "."))
            val = clean_amount(m)
            print(f"Clean amount: {val}")
            # If leading digit is '2' and stripping it gives a value close to claimed_amount, use corrected value
            if claimed_amount < val:
                str_int = str(int(val))
                if str_int.startswith("2") and len(str_int) > 1:
                    decimal_part = round((val % 1) * 100)
                    corrected = float(f"{str_int[1:]}.{decimal_part:02d}")
                    if abs(corrected - claimed_amount) < 20:
                        # print(f"[{key}] OCR ₹->2 fix applied: {val} -> {corrected}")
                        amounts.append(corrected)
                        continue

            if claimed_amount > val * 5:
                continue
            
            amounts.append(val)
        except:
            pass
    
    if claimed_amount is not None:
        amounts = sorted(amounts, key=lambda x: abs(x - claimed_amount))

    return amounts, corrected_gst, gst_was_corrected


def tamper_detection(path):

    original = cv2.imread(path)
    temp = "compressed.jpg"
    cv2.imwrite(temp, original, [cv2.IMWRITE_JPEG_QUALITY, 90])
    compressed = cv2.imread(temp)
    diff = cv2.absdiff(original, compressed)
    score = np.mean(diff)
    # print(f"Tamper score:{score}")
    return score

#  <-- Helper Functions to check Product Listed in Receipt appears in Proof Images -->
# def extract_products(text):
#                                                 # remove it
#     products = []
#     lines = text.split("\n")

#     for line in lines:

#         line = line.strip()
#         # match item number at beginning
#         match = re.match(r"^\d+[\.:,_\- ]+([A-Z][A-Z\s]{3,})", line)
#         if match:        
#             product = match.group(1).strip()
#             # remove bracket info
#             product = re.sub(r"\(.*?\)", "", product)
#             # remove small garbage words
#             if len(product) > 4:
#                 products.append(product)

#     print(f"Products: {products}")

#     return products[:5]

# def product_in_image(image_path, product):                    Dont delete

#     image = Image.open(image_path).convert("RGB")             remove

#     prompts = [
#         f"a photo of {product}",
#         f"{product} on a table",
#         f"{product} product"
#     ]

#     inputs = processor(
#         text=prompts,
#         images=image,
#         return_tensors="pt",
#         padding=True
#     ).to(device)

#     outputs = model(**inputs)

#     image_embeds = outputs.image_embeds
#     text_embeds = outputs.text_embeds

#     sims = torch.cosine_similarity(image_embeds, text_embeds)

#     sim = sims.max().item()   # pick best similarity

#     print(f"Sim for {product}: {sim}")

#     return sim

# def verify_products(image_path, products):                        Dont delete

#     matches = []

#     for p in products:                    remove

#         similarity = product_in_image(image_path, p)

#         if similarity > 0.25:
#             matches.append(p)
    
#     print(f"Matches: {matches}")
#     return matches

# def classify_crop(crop, products):
#                                                     # remove
#     pil = Image.fromarray(crop)

#     inputs = processor(
#         text=products,
#         images=pil,
#         return_tensors="pt",
#         padding=True
#     ).to(device)

#     outputs = model(**inputs)

#     sims = torch.cosine_similarity(
#         outputs.image_embeds,
#         outputs.text_embeds
#     )

#     best = sims.argmax().item()

#     return products[best], sims[best].item()

# ─────────────────────────────────────────────
#  Endpoints
# ─────────────────────────────────────────────

@app.get("/health")
def health():
    return {"status": "ok"}


@app.post("/analyze")
def analyze(data: AnalyzeRequest):
    """
    Scoring breakdown (max = 100):
      - Image quality / blur check :  20 pts
      - Duplicate Image             : 20 pts
      - Semantic similarity         :  60 pts
     
 
    Status rules:
      - score >= 70  AND no hard failures  →  "verified"
      - score >= 40  OR  any hard failure  →  "review"
      - score <  40                        →  "rejected"
 
    Hard failures (cap status at "review" even if score >= 70):
      - cause_mismatch  (semantic similarity below threshold)
      - duplicate_image
    """

    if not os.path.exists(data.image_path):
        return {"error": "Image not found"}

    flags = []
    score = 0
    hard_failures = []

    # ── 1. Image quality — blur detection (max 20 pts) ──────────────────────
    img = cv2.imread(data.image_path)
    gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)

    blur_score = cv2.Laplacian(gray, cv2.CV_64F).var()
    print(f"Blur score: {blur_score}")

    if blur_score < 50:
        flags.append("blurry_image")
    else:
        score += 20

    # ── 2. Perceptual hash — duplicate detection ─────────────────────────────
    # try to detect diplicate image after retrieving previous phashes from datatbase. Then add 10 pts. and make blur to 10pts
    pil_img = Image.open(data.image_path).convert("RGB")

    current_phash = imagehash.phash(pil_img)

    for prev in data.previous_phashes:
        prev_hash = imagehash.hex_to_hash(prev)

        # Hamming distance
        distance = current_phash - prev_hash

        if distance < 5:
            flags.append("Duplicate image")
            hard_failures.append("duplicate_image")
            break
    
    # After verified images, store the hash in your database.Store: Cause ID, Image URL, Perceptual Hash, CLIP Embedding
    data.previous_phashes.append(str(current_phash))
    print(f"Previous phashes: {data.previous_phashes}")
    score += 20

    # ── 3. Semantic similarity — cause vs image (max 60 pts) ─────────────────
    CLIP_LOW  = 0.15   # score = 0  below this
    CLIP_HIGH = 0.35   # score = 60 at or above this
    CLIP_THRESHOLD = 0.20   # below → hard cause_mismatch flag

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
    print(f"Raw CLIP similarity: {similarity:.4f}")

    # semantic_score = int((similarity + 1) / 2 * 80)  # normalize -1..1 → 0..80
    
    if similarity < CLIP_THRESHOLD:
        flags.append("cause_mismatch")
        hard_failures.append("cause_mismatch")
        semantic_pts = 0
    else:
        # linear map [CLIP_LOW, CLIP_HIGH] → [0, 60]
        clamped = max(CLIP_LOW, min(CLIP_HIGH, similarity))
        semantic_pts = int(((clamped - CLIP_LOW) / (CLIP_HIGH - CLIP_LOW)) * 60)
 
    score += semantic_pts
    print(f"Semantic pts: {semantic_pts} / 60  (similarity={similarity:.4f})")


    # ── 4. Product detection via YOLO + CLIP (max 40 pts) ────────────────────
    # product_matches = []
    #                                         # remove 
    # if data.products:
    #     results = yolo_model(data.image_path, conf=0.1)
    #     boxes = []
    #     for r in results:
    #         for box in r.boxes:
    #             x1, y1, x2, y2 = map(int, box.xyxy[0])
    #             boxes.append((x1, y1, x2, y2))
 
    #     raw_img = cv2.imread(data.image_path)
    #     matched = set()
 
    #     for (x1, y1, x2, y2) in boxes:
    #         crop = raw_img[y1:y2, x1:x2]
    #         label, sim_score = classify_crop(crop, data.products)
    #         print(f"Crop label: {label}, sim: {sim_score:.4f}")
    #         if sim_score > 0.2:
    #             matched.add(label)
 
    #     product_matches = list(matched)
    #     ratio = len(matched) / len(data.products)
 
    #     if ratio >= 0.6:
    #         score += 40
    #     elif ratio >= 0.3:
    #         score += 20
    #         flags.append("partial_product_match")
    #     else:
    #         score += 0
    #         flags.append("products_not_visible")



    # ── 5. Final decision ────────────────────────────────────────────────────    
    # Cap score at 100
    score = min(score, 100)
 
    if hard_failures:
        # One or more critical checks failed — cannot be "verified"
        status = "review" if score >= 40 else "rejected"
    elif score >= 70:
        status = "verified"
    elif score >= 40:
        status = "review"
    else:
        status = "rejected"
 
    return {
        "final_score": score,
        "validation_status": status,
        "semantic_similarity": round(similarity, 4),
        # "product_matches": product_matches,
        "hard_failures": hard_failures,   # surfaces what blocked "verified"
        "flags": flags
    }


# <-- analyze-reciept Endpoint -->
@app.post("/analyze-receipt")
def analyze_receipt(data: ReceiptRequest):
    """
    Scoring breakdown (max = 100):
      - GST valid                   :  25 pts
      - No tampering detected       :  25 pts
      - Amount matches claimed      :  25 pts
      - Vendor reputation (GST ok)  :  25 pts
 
    Status rules:
      - score >= 70  AND no hard failures  →  "verified"
      - score >= 40  OR  any hard failure  →  "review"
      - score <  40                        →  "rejected"
 
    Hard failures (cap status at "review"):
      - invalid_gst
      - possible_tampering
      - amount_mismatch
    """

    valid, error = validate_file(data.receipt_path)

    if not valid:
        return {"error": error}

    flags = []
    score = 0
    hard_failures = []

    # convert pdf if needed
    path = convert_pdf(data.receipt_path)

    # Ensure jpg format 
    path = ensure_jpg(path)

    # preprocessing
    processed = preprocess_receipt(path)

    # OCR
    text = extract_receipt_text(processed)
    print(f"Extracted OCR Text:{text}")

    # Normalize text
    text = normalize_ocr_text(text)
    print(f"Normalize OCR Text:{text}")

    # parse fields
    amounts, gst, gst_corrected = parse_receipt(text, data.claimed_amount)
    print(f"Amount:{amounts}, GST: {gst}")
    # amount, gst, gst_corrected = parse_receipt(text, data.claimed_amount)
    # print(f"Amount:{amount}, GST: {gst}")

    # ── 1. GST validation (25 pts) ───────────────────────────────────────────
    gst_valid = validate_gst(gst)

    if gst_valid:
        score += 25
        if gst_corrected:
            flags.append("gst_ocr_corrected")  # soft flag, not penalized
    else:
        flags.append("invalid_gst")
        hard_failures.append("invalid_gst")


    # ── 2. Tamper detection (25 pts) ─────────────────────────────────────────
    tamper_score = tamper_detection(path)

    if tamper_score < 10:
        score += 25
    else:
        flags.append("possible_tampering")
        hard_failures.append("possible_tampering")

    # ── 3. Amount matching (25 pts) ──────────────────────────────────────────
    amount_found = False
    for amt in amounts:

        if abs(amt - data.claimed_amount) < 10:
            score += 25
            amount_found = True
            break

    if not amount_found:
        flags.append("amount_mismatch")
        hard_failures.append("amount_mismatch")
    # for amt in amounts:    

    #     if abs(amt - data.claimed_amount) < 10:
    #         score += 25
    #     else:
    #         flags.append("amount_mismatch")
    #         hard_failures.append("amount_mismatch")

    # else:
    #     flags.append("amount_not_found")
    #     hard_failures.append("amount_not_found")
    
    # if amount is not None:

    #     if abs(amount - data.claimed_amount) < 10:
    #         score += 25
    #     else:
    #         flags.append("amount_mismatch")
    #         hard_failures.append("amount_mismatch")

    # else:
    #     flags.append("amount_not_found")
    #     hard_failures.append("amount_not_found")

    # ── 4. Vendor reputation via GST (25 pts) ────────────────────────────────
    if gst_valid:
        score += 25


    # remove 
     # ── 5. Extract products (informational) ──────────────────────────────────
    # products = extract_products(text)           
    
    # ── 6. Final decision ────────────────────────────────────────────────────
    score = min(score, 100)
 
    if hard_failures:
        status = "review" if score >= 40 else "rejected"
    elif score >= 70:
        status = "verified"
    elif score >= 40:
        status = "review"
    else:
        status = "rejected"
 
    print(f"Score: {score}, Status: {status}")
 
    return {
        "detected_amount": amounts[0],
        "gst_number": gst,
        "gst_ocr_corrected": gst_corrected,
        "tamper_score": tamper_score,
        "receipt_score": score,
        "status": status,
        # "products": products,
        "hard_failures": hard_failures,   # surfaces what blocked "verified"
        "flags": flags
    }

