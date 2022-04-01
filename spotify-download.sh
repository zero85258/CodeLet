# EX: $1=https://open.spotify.com/album/2gxNz8ZRquL5UEQmZ5ObOl

rm -rf ~/spotdl
spotdl -b $1 --write-to /tmp/spotdl-tmp.txt #getAlbumList
mkdir ~/spotdl
for n in `cat /tmp/spotdl-tmp.txt`; do
  spotdl -f ~/spotdl -s $n  #downloadSongs
done
