# 下載bin 賦予執行權限,即可使用
apt-get update && apt-get install -y libnss3-tools wget
wget https://github.com/FiloSottile/mkcert/releases/download/v1.1.2/mkcert-v1.1.2-linux-amd64
chmod +x mkcert-v1.1.2-linux-amd64
mv mkcert-v1.1.2-linux-amd64 /usr/local/bin/mkcert

