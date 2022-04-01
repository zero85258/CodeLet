package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    rows := map[string]int{'stitch':626, 'mimi':627}
    for key, value := range rows {
	fmt.Fprintf(w, key value)
    }
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
