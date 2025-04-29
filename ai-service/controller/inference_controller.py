import torch
from flask import Flask, request, jsonify
from transformers import WhisperProcessor, WhisperForConditionalGeneration
import torchaudio
import numpy as np
import shutil
from concurrent.futures import ThreadPoolExecutor
from config.hf_config import hf_read_login

hf_read_login()

# Load mô hình và processor
# checkpoint_path = "./whisper-small-vn-streaming/checkpoint-step-2300"
checkpoint_path = "DoanNgocHieu/think_flow" 
processor = WhisperProcessor.from_pretrained(checkpoint_path)
model = WhisperForConditionalGeneration.from_pretrained(checkpoint_path)

device = "cuda" if torch.cuda.is_available() else "cpu"
print(f"Using device: {device}")
model.to(device)

def process_segment(segment, processor, model, device):
    input_features = processor(segment, sampling_rate=16000, return_tensors="pt").input_features.to(device)
    with torch.no_grad():
        predicted_ids = model.generate(input_features)
    transcript = processor.batch_decode(predicted_ids, skip_special_tokens=True)[0]
    return transcript

def transcribe_audio_parallel(audio_path, segment_duration=15, max_workers=5):
    audio, sr = torchaudio.load(audio_path)

    if sr != 16000:
        resampler = torchaudio.transforms.Resample(orig_freq=sr, new_freq=16000)
        audio = resampler(audio)

    audio = audio.mean(dim=0)  # Convert to mono
    audio = audio.numpy()

    segment_samples = segment_duration * 16000
    total_samples = audio.shape[0]

    segments = [
        audio[start:min(start + segment_samples, total_samples)]
        for start in range(0, total_samples, segment_samples)
    ]

    transcripts = []
    with ThreadPoolExecutor(max_workers=max_workers) as executor:
        futures = [executor.submit(process_segment, segment, processor, model, device) for segment in segments]
        for future in futures:
            transcripts.append(future.result())

    return " ".join(transcripts)

# if __name__ == "__main__":
#     # Test transcribe
#     test_audio_path = "D:\\Audio_train\\audio_giong_hue_1.mp3"  # Đường dẫn tới file audio để test
#     transcripts = transcribe_audio_parallel(test_audio_path)
#     print("Transcription result:")
#     print(transcripts)




