import os
import torch
from huggingface_hub import login
from datasets import load_dataset
from transformers import WhisperProcessor, WhisperForConditionalGeneration
from torch.utils.data import DataLoader, IterableDataset
from dataclasses import dataclass
from typing import List, Dict, Union
from config.hf_config import hf_read_login

# Login hugging face
hf_read_login()
print('We are logged in to Hugging Face now!')

# 2. Load dữ liệu dạng stream
train_dataset = load_dataset(
    "baohuynhbk14/vietnamese-speech-to-text-preprocessed-whisper-medium", 
    split="train", 
    streaming=True
)
test_dataset = load_dataset(
    "baohuynhbk14/vietnamese-speech-to-text-preprocessed-whisper-medium", 
    split="test", 
    streaming=True
)

# 3. Processor và model
processor = WhisperProcessor.from_pretrained("openai/whisper-small", language="vietnamese", task="transcribe")
model = WhisperForConditionalGeneration.from_pretrained("openai/whisper-small")
model.config.forced_decoder_ids = None
model.config.suppress_tokens = []
model.config.use_cache = False
model.to("cuda" if torch.cuda.is_available() else "cpu")

# 4. Data Collator
@dataclass
class DataCollatorSeq2SeqWithPadding:
    processor: any

    def __call__(self, features: List[Dict[str, Union[List[int], torch.Tensor]]]) -> Dict[str, torch.Tensor]:
        input_features = [{"input_features": feature["input_features"]} for feature in features]
        batch = self.processor.feature_extractor.pad(input_features, return_tensors="pt")

        label_features = [{"input_ids": feature["labels"]} for feature in features]
        labels_batch = self.processor.tokenizer.pad(label_features, return_tensors="pt")
        labels = labels_batch["input_ids"].masked_fill(labels_batch.attention_mask.ne(1), -100)

        if (labels[:, 0] == self.processor.tokenizer.bos_token_id).all().cpu().item():
            labels = labels[:, 1:]

        batch["labels"] = labels
        return batch

data_collator = DataCollatorSeq2SeqWithPadding(processor)

# 5. Streaming DataLoader
class HFIterableDataset(torch.utils.data.IterableDataset):
    def __init__(self, dataset):
        self.dataset = dataset

    def __iter__(self):
        return iter(self.dataset)

train_iterable = HFIterableDataset(train_dataset)
train_dataloader = DataLoader(train_iterable, batch_size=4, collate_fn=data_collator)

# 6. Optimizer
optimizer = torch.optim.AdamW(model.parameters(), lr=3e-5)

# 7. Vòng lặp training thủ công
device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
model.train()
save_dir = "./whisper-small-vn-streaming"
os.makedirs(save_dir, exist_ok=True)

step = 600
checkpoint_interval = 100
max_steps = 5000  # dừng sau bao nhiêu bước tùy bạn

for batch in train_dataloader:
    input_features = batch["input_features"].to(device)
    labels = batch["labels"].to(device)

    outputs = model(input_features=input_features, labels=labels)
    loss = outputs.loss

    loss.backward()
    optimizer.step()
    optimizer.zero_grad()

    step += 1
    print(f"🧪 Step {step} | Loss: {loss.item():.4f}")

    if step % checkpoint_interval == 0:
        ckpt_path = os.path.join(save_dir, f"checkpoint-step-{step}")
        model.save_pretrained(ckpt_path)
        processor.save_pretrained(ckpt_path)
        print(f"💾 Đã lưu checkpoint tại bước {step}")

    if step >= max_steps:
        print("✅ Huấn luyện hoàn tất!")
        break
