from huggingface_hub import upload_folder, login
from config.hf_config import hf_write_login

hf_write_login()

upload_folder(
    repo_id="DoanNgocHieu/think_flow",
    folder_path="./whisper-small-vn-streaming/checkpoint-step-3000",
    commit_message="Ghi đè mô hình bằng checkpoint step 3000",
    allow_patterns="**",  
    delete_patterns="**"  
)
