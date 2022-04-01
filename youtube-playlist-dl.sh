# $1=url
youtube-dl -cit --extract-audio --audio-format mp3 $1
#ffmpeg -i input.webm -acodec libmp3lame -aq 4 output.mp3
