${1:-/dev/stdin} | xargs -I {} curl -s {} | w3m -T text/html -dump
