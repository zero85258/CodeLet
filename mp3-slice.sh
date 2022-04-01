youtube-dl $1
ffmpeg -i file.mkv -ss 00:00:20 -to 00:00:40 -c copy file-2.mkv
ffmpeg -i your_audio_file.mp3 -acodec copy -t 00:30:00 -ss 00:00:00 half_hour_split.mp3
