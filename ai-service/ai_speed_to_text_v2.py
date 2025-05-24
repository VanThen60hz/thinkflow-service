from huggingface_hub import login
from datasets import load_dataset, DatasetDict
from transformers import (
    WhisperForConditionalGeneration,
    WhisperProcessor,
    Seq2SeqTrainingArguments,
    Seq2SeqTrainer,
)
import os
import evaluate
import torch
from dataclasses import dataclass
from typing import Dict, List, Union
from config.hf_config import hf_read_login

# Login hugging face
hf_read_login()
print('We are logged in to Hugging Face now!')

# 2. Load dataset đã tiền xử lý sẵn (gồm input_features & labels)
common_voice = DatasetDict()
common_voice["train"] = load_dataset(
    "baohuynhbk14/vietnamese-speech-to-text-preprocessed-whisper-medium", 
    split="train", 
    streaming=True
)
common_voice["test"] = load_dataset(
    "baohuynhbk14/vietnamese-speech-to-text-preprocessed-whisper-medium", 
    split="test", 
    streaming=True
)

# 3. Load processor và model
processor = WhisperProcessor.from_pretrained("openai/whisper-small", language="vietnamese", task="transcribe")
model = WhisperForConditionalGeneration.from_pretrained("openai/whisper-small")
model.config.forced_decoder_ids = None
model.config.suppress_tokens = []
model.config.use_cache = False

# 4. Định nghĩa Data Collator
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

data_collator = DataCollatorSeq2SeqWithPadding(processor=processor)

# 5. Định nghĩa metric
metric = evaluate.load("wer")

def compute_metrics(pred):
    pred_ids = pred.predictions
    label_ids = pred.label_ids
    label_ids[label_ids == -100] = processor.tokenizer.pad_token_id

    pred_str = processor.tokenizer.batch_decode(pred_ids, skip_special_tokens=True)
    label_str = processor.tokenizer.batch_decode(label_ids, skip_special_tokens=True)

    wer = 100 * metric.compute(predictions=pred_str, references=label_str)
    return {"wer": wer}

# 6. Training arguments
training_args = Seq2SeqTrainingArguments(
    output_dir="./whisper-small-vn-v2",
    per_device_train_batch_size=16,
    gradient_accumulation_steps=2,
    learning_rate=3e-6,
    warmup_steps=500,
    num_train_epochs=3,
    gradient_checkpointing=True,
    fp16=True,
    eval_strategy="steps",
    per_device_eval_batch_size=8,
    predict_with_generate=True,
    generation_max_length=225,
    save_steps=500,
    save_total_limit=3,
    eval_steps=500,
    logging_steps=25,
    report_to=["tensorboard"],
    generation_num_beams=5,
    load_best_model_at_end=True,
    metric_for_best_model="wer",
    greater_is_better=False,
    push_to_hub=False
)

# 7. Tạo trainer và train
trainer = Seq2SeqTrainer(
    args=training_args,
    model=model,
    train_dataset=common_voice["train"],
    eval_dataset=common_voice["test"],
    data_collator=data_collator,
    compute_metrics=compute_metrics,
    tokenizer=processor.feature_extractor,
)

# Kiểm tra checkpoint [add]
checkpoint_path = "./whisper-small-vn-v2/checkpoint-3"
resume_training = os.path.exists(checkpoint_path) 

# Train model [add]
if resume_training:
    print("Tiếp tục training từ checkpoint...")
    trainer.train(resume_from_checkpoint=checkpoint_path)
else:
    print("Training mới từ đầu...")
    trainer.train()

# trainer.train()
processor.save_pretrained(checkpoint_path)
print("Tokenizer đã được lưu vào checkpoint.")
