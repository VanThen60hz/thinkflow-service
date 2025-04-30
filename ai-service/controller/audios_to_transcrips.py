from controller.inference_controller import transcribe_audio_parallel 
from services.openai_utils import correct_transcrip

def transcribe_multiple_audios(audio_paths, segment_duration=15, max_workers=5):
    transcripts = []
    for audio_path in audio_paths:
        transcript = transcribe_audio_parallel(audio_path, segment_duration, max_workers)
        correct_transcrip_result = correct_transcrip(transcript)
        transcripts.append(correct_transcrip_result)
    return transcripts

def transcribe_from_an_audio(audio_path, segment_duration=15, max_workers=5):
    correct_transcrip_result = ''
    transcript = transcribe_audio_parallel(audio_path, segment_duration, max_workers)
    correct_transcrip_result = correct_transcrip(transcript)
    return correct_transcrip_result

def transcribe_from_an_audio_test(audio_path, segment_duration=15, max_workers=5):
    correct_transcrip_result = ''
    transcript = transcribe_audio_parallel(audio_path, segment_duration, max_workers)
    correct_transcrip_result = transcript
    return correct_transcrip_result
    
