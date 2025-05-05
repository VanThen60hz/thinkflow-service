import json
import os
import time
import tempfile
from flask import Blueprint, request, jsonify
from controller.text_to_mindmap_controller import mindmap_and_structure 
from services.openai_utils import summarize_to_text_model 
from controller.inference_controller import transcribe_audio_parallel
from controller.audios_to_transcrips import transcribe_multiple_audios, transcribe_from_an_audio, transcribe_from_an_audio_test

transcript_routes = Blueprint('transcript_routes', __name__)

# -------------------------------------------------------------
# ----------------------------Transcript-----------------------
# -------------------------------------------------------------
# 
# Đưa ra list transcript từ list audio
# 
@transcript_routes.route("/list_transcripts_from_audios", methods=["POST"])
def list_transcripts_from_audios():
    start_time = time.time()

    if 'files' not in request.files:
        return jsonify({"error": "No files part"}), 400

    files = request.files.getlist('files')
    if len(files) == 0:
        return jsonify({"error": "No selected files"}), 400

    tmp_paths = []
    try:
        # Lưu tạm tất cả các file
        for file in files:
            suffix = os.path.splitext(file.filename)[-1] or ".tmp"
            with tempfile.NamedTemporaryFile(delete=False, suffix=suffix) as tmp:
                file.save(tmp.name)
                tmp_paths.append(tmp.name)

        # Transcribe tất cả audio
        transcripts = transcribe_multiple_audios(tmp_paths)

    finally:
        # Xóa file tạm
        for path in tmp_paths:
            if os.path.exists(path):
                os.remove(path)

    elapsed_time = time.time() - start_time
    print(f"⏱️ Thời gian hoàn thành: {elapsed_time:.2f} giây")

    return jsonify({
        "list_transcripts": transcripts,
        "time_taken_sec": round(elapsed_time, 2)
    })

# 
# Đưa ra transcript từ một audio
# 
@transcript_routes.route("/transcript_from_an_audio", methods=["POST"])
def transcript_from_an_audio():
    start_time = time.time()

    file = request.files.get("file")
    if file is None:
        return jsonify({"error": "No file uploaded"}), 400

    try:
        # Lưu tạm 1 các file
        suffix = os.path.splitext(file.filename)[-1] or ".tmp"
        with tempfile.NamedTemporaryFile(delete=False, suffix=suffix) as tmp:
            file.save(tmp.name)
            tmp_path = tmp.name

        # Transcribe 1 audio
        transcript = transcribe_from_an_audio(tmp_path)

    finally:
        # Xóa file tạm
        if os.path.exists(tmp_path):
            os.remove(tmp_path)

    elapsed_time = time.time() - start_time
    print(f"⏱️ Thời gian hoàn thành: {elapsed_time:.2f} giây")

    return jsonify({
        "transcript": transcript,
        "time_taken_sec": round(elapsed_time, 2)
    })

# 
# Đưa ra transcript từ một audio (ban test) 
# 
@transcript_routes.route("/transcript_from_an_audio_test", methods=["POST"])
def transcript_from_an_audio_test():
    start_time = time.time()

    file = request.files.get("file")
    if file is None:
        return jsonify({"error": "No file uploaded"}), 400

    try:
        # Lưu tạm 1 các file
        suffix = os.path.splitext(file.filename)[-1] or ".tmp"
        with tempfile.NamedTemporaryFile(delete=False, suffix=suffix) as tmp:
            file.save(tmp.name)
            tmp_path = tmp.name

        # Transcribe 1 audio
        transcript = transcribe_from_an_audio_test(tmp_path)

    finally:
        # Xóa file tạm
        if os.path.exists(tmp_path):
            os.remove(tmp_path)

    elapsed_time = time.time() - start_time
    print(f"⏱️ Thời gian hoàn thành: {elapsed_time:.2f} giây")

    return jsonify({
        "transcript": transcript,
        "time_taken_sec": round(elapsed_time, 2)
    })





