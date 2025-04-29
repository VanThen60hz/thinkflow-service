from huggingface_hub import upload_folder, login
from config.hf_config import hf_write_login

hf_write_login()

upload_folder(
    repo_id="DoanNgocHieu/think_flow",
    folder_path="./whisper-small-vn-streaming/checkpoint-step-2300",   
    commit_message="Upload model lần đầu"
)
