DATE_LEN=$(echo $1 | wc -c | awk '{print $1}')

if [ ${DATE_LEN} -eq '7' ];then
  aws s3 ls --recursive --human-readable s3://XXXdata/sparkmonthdata/DMP/$1/
elif [ ${DATE_LEN} -eq '9' ];then
  aws s3 ls --human-readable s3://XXXdata/flat_data/DMP/$1/
  echo "----------------------------------------------------------------";
  aws s3 ls --recursive --human-readable s3://XXXdata/sparktempdata/DMP/$1/
else
  echo "gag"
fi
