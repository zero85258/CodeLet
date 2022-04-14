package main

import (
	// "html/template"
	"github.com/gin-gonic/gin"
	"net/http"
)

// func tmpl(w http.ResponseWriter, r *http.Request) {
// 	t1, err := template.ParseFiles("tpl.html")
// 	if err != nil {
// 		panic(err)
// 	}
// 	t1.Execute(w, "hello world")
// }

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tpl.html", gin.H{
			"title": "hello gin " + " method",
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// C.HTML(http.StatusOK, "tpl.html", gin.H{
	// 	"title": "hello gin " + " method",
	// })

	// server := http.Server{
	// 	Addr: "127.0.0.1:8080",
	// }
	// http.HandleFunc("/tmpl", tmpl)
	// server.ListenAndServe()
}
