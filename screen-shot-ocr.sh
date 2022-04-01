# $1=指定哪個螢幕,1是主螢幕,2是副螢幕
screencapture -- -D $1 /tmp/sc.jpg && touch /tmp/ocr.txt && tesseract -l eng /tmp/sc.jpg /tmp/ocr.txt && less /tmp/ocr.txt
