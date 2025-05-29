import openai
import json
import os
from services.prompts import mindmap_prompt, summary_prompt, audio_prompt
from dotenv import load_dotenv

load_dotenv() 

openai.api_key = os.getenv("OPENAI_API_KEY")

def mindmap_to_json_model(text):
    response = openai.chat.completions.create(
        model="gpt-4-turbo",
        messages=[
            {"role": "system", "content": mindmap_prompt}, 
            {"role": "user", "content": f"Văn bản cần thực hiện theo yêu cầu:\n\"\"\"{text}\"\"\""}      
        ]
    )
    ai_reply = response.choices[0].message.content.strip()
    print(ai_reply)
    try:
        return json.loads(ai_reply)
    except json.JSONDecodeError:
        return None

def summarize_to_text_model(text):
    response = openai.chat.completions.create(
        model="gpt-4o",
        messages=[
            {"role": "system", "content": summary_prompt},
            {"role": "user", "content": f"Văn bản cần tóm tắt theo yêu cầu:\n\"\"\"{text}\"\"\""}
        ]

    )
    return response.choices[0].message.content.strip()

def correct_transcrip(text):
    response = openai.chat.completions.create(
        model="gpt-3.5-turbo",
        messages=[
            {"role": "system", "content": audio_prompt}, 
            {"role": "user", "content": f"Văn bản cần thực hiện theo yêu cầu:\n\"\"\"{text}\"\"\""}    
        ]
    )
    return response.choices[0].message.content.strip()
