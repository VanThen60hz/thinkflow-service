import openai
import json
import os
from services.prompts import mindmap_prompt, summary_prompt, audio_prompt
from dotenv import load_dotenv

load_dotenv() 

openai.api_key = os.getenv("OPENAI_API_KEY")

def mindmap_to_json_model(text):
    full_prompt = mindmap_prompt.format(text=text)
    response = openai.chat.completions.create(
        model="gpt-3.5-turbo",
        messages=[{"role": "user", "content": full_prompt}]
    )
    ai_reply = response.choices[0].message.content.strip()
    try:
        return json.loads(ai_reply)
    except json.JSONDecodeError:
        return None

def summarize_to_text_model(text):
    full_prompt = summary_prompt.format(text=text)
    response = openai.chat.completions.create(
        model="gpt-3.5-turbo",
        messages=[{"role": "user", "content": full_prompt}]
    )
    return response.choices[0].message.content.strip()

def correct_transcrip(text):
    full_prompt = audio_prompt.format(text=text)
    response = openai.chat.completions.create(
        model="gpt-3.5-turbo",
        messages=[{"role": "user", "content": full_prompt}]
    )
    return response.choices[0].message.content.strip()
