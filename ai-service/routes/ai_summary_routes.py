import json
import os
import time
import tempfile
from flask import Blueprint, request, jsonify
from controller.text_to_mindmap_controller import mindmap_and_structure 
from services.openai_utils import summarize_to_text_model 
from controller.inference_controller import transcribe_audio_parallel
from controller.audios_to_transcrips import transcribe_multiple_audios, transcribe_from_an_audio, transcribe_from_an_audio_test

summary_routes = Blueprint('summary_routes', __name__)

# ----------------------------------------------------------
# ---------------------------Summary------------------------
# ----------------------------------------------------------
# 
# Summary from text
# 
@summary_routes.route('/summary_from_text', methods=['POST'])
def summary_from_text():
    try:
        data = request.get_json()
        if not data or "text" not in data:
            return jsonify({
                "status": 400,
                "message": "Thiếu trường 'text' trong request body."
            }), 400

        summary_result = summarize_to_text_model(data["text"])
        return jsonify({"summary": summary_result})
    except json.JSONDecodeError:
        return jsonify({
           "status": 500,
           "message": "Không thể phân tích phản hồi từ AI."
        }), 500

# 
# Đưa ra summary từ tất cả audio
# 
@summary_routes.route("/summary_from_all_audios", methods=["POST"])
def summary_from_all_audios():
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
        
        full_text = " ".join(transcripts)

    finally:
        # Xóa file tạm
        for path in tmp_paths:
            if os.path.exists(path):
                os.remove(path)

    # Gọi AI để tóm tắt văn bản
    summary_result = summarize_to_text_model(full_text)

    elapsed_time = time.time() - start_time
    print(f"⏱️ Thời gian hoàn thành: {elapsed_time:.2f} giây, {full_text}")

    return jsonify({"summary_from_audios": summary_result, "time_taken_sec": round(elapsed_time, 2)})


# 
# Summary từ all transcript audios input
# 
@summary_routes.route('/summary_from_all_transcript_audios', methods=['POST'])
def summary_from_all_transcript_audios():
    data = request.get_json()

    if not data:
        return jsonify({
            "status": 400,
            "message": "Thiếu dữ liệu trong request body."
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

    try:
        summary_result = summarize_to_text_model(all_text_audio)
        return jsonify({"summary_all_transcript_audios": summary_result})
    except Exception:
        return jsonify({
            "status": 500,
            "message": "Không thể phân tích phản hồi từ AI."
        }), 500
    
# 
#  Summary từ tất cả transcrip audios và text lớn 
# 
@summary_routes.route('/summary_note_from_all_text', methods=['POST'])
def summary_note_from_all_text():
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

    if not full_text:
        return jsonify({
            "status": 400,
            "message": "Không có nội dung để xử lý (cả text và audio đều trống)."
        }), 400

    try:
        summary_result = summarize_to_text_model(full_text)
        return jsonify({"summary_from_all_text": summary_result})
    except Exception:
        return jsonify({
            "status": 500,
            "message": "Không thể phân tích phản hồi từ AI."
        }), 500
    

# 
#  Summary từ tất cả audios và text
# 
@summary_routes.route("/summary_note_from_all_audio_and_text", methods=["POST"])
def summary_note_from_all_audio_and_text():
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

    try:
        summary_result = summarize_to_text_model(full_text)
        return jsonify({"summary_audios_and_text": summary_result})
    except Exception:
        return jsonify({
            "status": 500,
            "message": "Không thể phân tích phản hồi từ AI."
        }), 500