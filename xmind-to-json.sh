xmindparser $1 -json > /dev/null 2>&1 
noExtPath=$(echo $1| cut -d'.' -f 1)
cat $noExtPath.json | ascii2uni -a U -q
rm $noExtPath.json
