from fastapi import FastAPI
from pydantic import BaseModel
import cv2
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

class AnalyzeRequest(BaseModel):
    image_path: str
    cause_text: str
    previous_phashes: list[str] = []


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

    # -------- CLIP Semantic Similarity --------
    # image = Image.open(data.image_path).convert("RGB")

    # inputs = processor(
    #     text=[data.cause_text],
    #     images=image,
    #     return_tensors="pt",
    #     padding=True
    # ).to(device)

    # outputs = model(**inputs)

    # image_embeds = outputs.image_embeds
    # text_embeds = outputs.text_embeds

    # # normalize embeddings
    # image_embeds = image_embeds / image_embeds.norm(dim=-1, keepdim=True)
    # text_embeds = text_embeds / text_embeds.norm(dim=-1, keepdim=True)

    # similarity = (image_embeds @ text_embeds.T).item()

    # semantic_score = int((similarity + 1) / 2 * 80)
    # print(f"Semantic similarity: {similarity:.4f}, score: {semantic_score}")
    # score += semantic_score

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