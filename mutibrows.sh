# 一口氣導覽多個網頁
pbpaste | xargs -I {} curl -s {} | w3m -T text/html -dump | less
