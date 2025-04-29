import openai
import gradio as gr
from PIL import Image
import requests
from io import BytesIO
import os
from dotenv import load_dotenv

load_dotenv() 

openai.api_key = os.getenv("OPENAI_API_KEY")

def generate_image_from_text(prompt):
    response = openai.images.generate(
        model="dall-e-3",  
        prompt=prompt,
        size="1024x1024",
        quality="standard",
        n=1,
    )
    image_url = response.data[0].url
    image_response = requests.get(image_url)
    image = Image.open(BytesIO(image_response.content))
    return image

# Tạo giao diện Gradio
demo = gr.Interface(
    fn=generate_image_from_text,
    inputs=gr.Textbox(label="Nhập mô tả hình ảnh"),
    outputs=gr.Image(type="pil", label="Ảnh được tạo"),
    title="Tạo Ảnh Từ Văn Bản",
    description="Nhập mô tả (ví dụ: 'một con vịt đang được quay') và hệ thống sẽ tạo ảnh cho bạn bằng DALL·E!"
)

demo.launch()

