import os
from huggingface_hub import login
from dotenv import load_dotenv

load_dotenv()

def hf_read_login():
    HF_READ_TOKEN = os.getenv("HUGGING_FACE_TOKEN")
    if HF_READ_TOKEN:
        login(token=HF_READ_TOKEN)
        print("Da login huggingface doc")

def hf_write_login():
    HF_WRITE_TOKEN = os.getenv("HUGGING_FACE_TOKEN_WRITE")
    if HF_WRITE_TOKEN:
        login(token=HF_WRITE_TOKEN)
        print("Da login huggingface viet")
    else:
        print("Không tìm thấy biến môi trường HUGGING_FACE_TOKEN_WRITE")
