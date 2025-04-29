from controller.inference_controller import transcribe_audio_parallel 
from services.openai_utils import correct_transcrip

def transcribe_multiple_audios(audio_paths, segment_duration=15, max_workers=5):
    transcripts = []
    for audio_path in audio_paths:
        transcript = transcribe_audio_parallel(audio_path, segment_duration, max_workers)
        correct_transcrip_result = correct_transcrip(transcript)
        transcripts.append(correct_transcrip_result)
    return transcripts

# if __name__ == "__main__":
#     # Test transcribe
#     audios_test = [
#                    "D:\\Audio_train\\audio_giong_ha_noi_1.mp3",
#                    "D:\\Audio_train\\audio_giong_ha_noi_2.mp3"
#                   ] 
#     transcripts = transcribe_multiple_audios(audios_test)
#     print("Transcription result:")
#     print(transcripts)