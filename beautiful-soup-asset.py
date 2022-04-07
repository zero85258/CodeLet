import requests
from bs4 import BeautifulSoup

url = "https://data.kcg.gov.tw/dataset/108kaohs"
for i in range(10): #往上爬3頁
    r = requests.get(url)
    soup = BeautifulSoup(r.text, "html.parser")
    sel = soup.select("#dataset-resources > ul")  # 標題
    u = soup.select("div.btn-group.btn-group-paging a") #a標籤
    print ("本頁的URL為"+url)
    url = "https://www.ptt.cc" + u[1]["href"] #上一頁的網址

for s in sel: #印出網址跟標題
    print("https://www.ptt.cc", s["href"], s.text)

# with open('test.csv','w',encoding="utf-8") as f:

# a4e91d42-89f4-4735-bd00-919260d4f4e1
# https://data.kcg.gov.tw/dataset/e7f76e4b-e9ab-468f-a0aa-6c5bc6251e5b/resource/a4e91d42-89f4-4735-bd00-919260d4f4e1/download/10804-2.csv

# with open('test.csv', 'w', encoding="utf-8") as f:
# writer = csv.writer(f)

#dataset-resources > ul

[
    {"title": "高雄市不動產買賣統計(108年1月)",
     "href": "https://data.kcg.gov.tw/dataset/e7f76e4b-e9ab-468f-a0aa-6c5bc6251e5b/resource/b3f3512b-168a-4389-a44a-5552f7bc6c44/download/10801-2.csv"},
    {"title": "高雄市不動產買賣統計(108年2月)",
     "href": "https://data.kcg.gov.tw/dataset/e7f76e4b-e9ab-468f-a0aa-6c5bc6251e5b/resource/9f491de5-ee3f-482c-b512-c860cdb05829/download/10802-1.csv"},
    {"title": "高雄市不動產買賣統計(108年3月)",
     "href": "https://data.kcg.gov.tw/dataset/e7f76e4b-e9ab-468f-a0aa-6c5bc6251e5b/resource/093d3125-f4da-4112-8417-c5dfbb12765b/download/10803-2.csv"},
    {"title": "高雄市不動產買賣統計(108年4月)",
     "href": "https://data.kcg.gov.tw/dataset/e7f76e4b-e9ab-468f-a0aa-6c5bc6251e5b/resource/a4e91d42-89f4-4735-bd00-919260d4f4e1/download/10804-2.csv"},
    {"title": "高雄市不動產買賣統計(108年5月)",
     "href": "https://data.kcg.gov.tw/dataset/e7f76e4b-e9ab-468f-a0aa-6c5bc6251e5b/resource/544363fa-335e-405a-99fc-1bbf71a1fb5c/download/10805.csv"},
    {"title": "高雄市不動產買賣統計(108年6月)",
     "href": "https://data.kcg.gov.tw/dataset/e7f76e4b-e9ab-468f-a0aa-6c5bc6251e5b/resource/6cbbbc4b-b219-4b70-aa4a-31622a1f3b97/download/10806-2.csv"},
    {"title": "高雄市不動產買賣統計(108年7月)",
     "href": "https://data.kcg.gov.tw/dataset/e7f76e4b-e9ab-468f-a0aa-6c5bc6251e5b/resource/1ab9811e-28ef-4c55-9ab4-241a4f524a44/download/10807-2.csv"},
    {"title": "高雄市不動產買賣統計(108年8月)",
     "href": "https://data.kcg.gov.tw/dataset/e7f76e4b-e9ab-468f-a0aa-6c5bc6251e5b/resource/65ece871-71e9-449a-a18e-b9c27b13ae53/download/10808-2.csv"},
    {"title": "高雄市不動產買賣統計(108年9月)",
     "href": "https://data.kcg.gov.tw/dataset/e7f76e4b-e9ab-468f-a0aa-6c5bc6251e5b/resource/9177b335-84bd-4330-929c-e7c7adbae7e5/download/10809-2.csv"},
    {"title": "高雄市不動產買賣統計(108年10月)",
     "href": "https://data.kcg.gov.tw/dataset/e7f76e4b-e9ab-468f-a0aa-6c5bc6251e5b/resource/27f4caa3-aa74-4653-8a92-263d6ec38652/download/10810-2.csv"},
    {"title": "高雄市不動產買賣統計(108年11月)",
     "href": "https://data.kcg.gov.tw/dataset/e7f76e4b-e9ab-468f-a0aa-6c5bc6251e5b/resource/44c2f8c6-c4cd-4615-b591-8f0f5172e322/download/10811-.csv"},
    {"title": "高雄市不動產買賣統計(108年12月)",
     "href": "https://data.kcg.gov.tw/dataset/e7f76e4b-e9ab-468f-a0aa-6c5bc6251e5b/resource/f5de9d98-bfae-4f87-a814-0b4378e7cee7/download/10812-1.csv"},
]


# Apache Sqoop
