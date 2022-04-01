ffmpeg -protocol_whitelist "file,http,https,tcp,tls" \
-i $1 \
-c copy "$2"
