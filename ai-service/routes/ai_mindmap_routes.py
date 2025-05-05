import json
import os
import time
import tempfile
from flask import Blueprint, request, jsonify
from controller.text_to_mindmap_controller import mindmap_and_structure 
from services.openai_utils import summarize_to_text_model 
from controller.inference_controller import transcribe_audio_parallel
from controller.audios_to_transcrips import transcribe_multiple_audios, transcribe_from_an_audio, transcribe_from_an_audio_test

mindmap_routes = Blueprint('mindmap_routes', __name__)

# -------------------------------------------------------------
# ----------------------------MindMap--------------------------
# -------------------------------------------------------------
# 
# Mindmap từ text và all text audio input
# 
@mindmap_routes.route('/mindmap_note_from_all_text', methods=['POST'])
def mindmap_note_from_all_text():
    data = request.get_json()

    if not data:
        return jsonify({
            "status": 400,
            "message": "Thiếu dữ liệu trong request body."
        }), 400

    if "text" not in data:
        return jsonify({
            "status": 400,
            "message": "Thiếu trường 'text' trong request body."
        }), 400

    if "list_transcrip_text" not in data:
        return jsonify({
            "status": 400,
            "message": "Thiếu trường 'list_transcrip_text' trong request body."
        }), 400

    if not isinstance(data["list_transcrip_text"], list):
        return jsonify({
            "status": 400,
            "message": "'list_transcrip_text' phải là dạng danh sách (list)."
        }), 400

    # Tập hợp tất cả các text
    all_text_audio = "".join(data["list_transcrip_text"])
    full_text = data["text"] + " " + all_text_audio

    try:
        result_json = mindmap_and_structure(full_text)
        result_dict = json.loads(result_json)
    except Exception:
        return jsonify({
            "status": 500,
            "message": "Không thể phân tích phản hồi từ AI."
        }), 500

    return jsonify(result_dict), result_dict.get("status", 200)

# 
# Mindmap từ tất cả audios và text
# 
@mindmap_routes.route("/mindmap_note_from_all_audio_and_text", methods=["POST"])
def mindmap_note_from_all_audio_and_text():
    start_time = time.time()
    
    # Lấy dữ liệu input
    files = request.files.getlist('files') if 'files' in request.files else []
    text = request.form.get('text', '')

    list_transcript = []

    # Nếu có file audio thì xử lý
    if files:
        for file in files:
            if file.filename == '':
                continue
            suffix = os.path.splitext(file.filename)[-1] or ".tmp"
            with tempfile.NamedTemporaryFile(delete=False, suffix=suffix) as tmp:
                file.save(tmp.name)
                tmp_path = tmp.name

            try:
                # transcrip từng audio và thêm vào list_transcript
                transcript = transcribe_audio_parallel(tmp_path)
                list_transcript.append(transcript)
            finally:
                os.remove(tmp_path)

    # Ghép text + transcript
    all_text_audio = " ".join(list_transcript)
    full_text = (text + " " + all_text_audio).strip()

    if not full_text:
        return jsonify({
            "status": 400,
            "message": "Không có nội dung để xử lý (cả text và audio đều trống)."
        }), 400

    # Gọi AI để tạo mindmap
    try:
        summary_result_json = mindmap_and_structure(full_text)
        summary_result_dict = json.loads(summary_result_json)
    except json.JSONDecodeError:
        return jsonify({
           "status": 500,
           "message": "Không thể phân tích phản hồi từ AI."
        }), 500

    elapsed_time = time.time() - start_time
    print(f"⏱️ Thời gian hoàn thành: {elapsed_time:.2f} giây")

    return jsonify(summary_result_dict) , summary_result_dict.get("status", 200)


# 
# Mindmap từ summary
# 
@mindmap_routes.route("/mindmap_note_from_summary", methods=["POST"])
def mindmap_note_from_summary():
    start_time = time.time()
    
    summary = request.json.get('summary', '')

    if not summary:
        return jsonify({
            "status": 400,
            "message": "Không có nội dung để xử lý."
        }), 400

    # Gọi AI để tạo mindmap
    try:
        summary_result_json = mindmap_and_structure(summary)
        summary_result_dict = json.loads(summary_result_json)
    except json.JSONDecodeError:
        return jsonify({
           "status": 500,
           "message": "Không thể phân tích phản hồi từ AI."
        }), 500

    elapsed_time = time.time() - start_time
    print(f"⏱️ Thời gian hoàn thành: {elapsed_time:.2f} giây")

    return jsonify(summary_result_dict) , summary_result_dict.get("status", 200)