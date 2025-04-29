import os
from huggingface_hub import login

def hf_read_login():
    HF_READ_TOKEN = os.getenv("HUGGING_FACE_TOKEN")
    if HF_READ_TOKEN:
        login(token=HF_READ_TOKEN)

def hf_write_login():
    HF_WRITE_TOKEN = os.getenv("HUGGING_FACE_TOKEN_WRITE")
    if HF_WRITE_TOKEN:
        login(token=HF_WRITE_TOKEN)
