curl  -H "Content-Type: application/json" -X PUT 'http://localhost:9200/twitter/user/winwu' -d '{ "name" : "小 穎穎" }'
curl  -H "Content-Type: application/json" -X GET 'http://localhost:9200/twitter/user/winwu'
curl -sS -H "Content-Type: application/json" -X GET 'http://localhost:9200/twitter/user/winwu' | jq '._source'
