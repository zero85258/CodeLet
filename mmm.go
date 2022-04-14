package main

import "fmt"
import "github.com/gofrs/uuid"
import "net/http"
import "github.com/gin-gonic/gin"
import "time"
import "sync"
import "math/rand"
import "go-test/services"

// 宣吿結構
type Rhythm struct {
	Groove string   `json:"groove"`
	Bpm    uint8    `json:"bpm"`
	Count  string   `json:"count"`
	Style  []string `json:"style"`
}

type SafeCounter struct {
	mux sync.Mutex
}

// 宣告方法
func (r *Rhythm) setBpm(speed uint8) uint8 {
	r.Bpm = speed
	return r.Bpm
}

func music() {
	// 賦值
	funkGroove := &Rhythm{
		"funk",
		120,
		"4/4",
		[]string{"breakbeat", "rudiment"},
	}

	// 執行方法
	returnSpeed := funkGroove.setBpm(115)
	fmt.Printf("%s\n", returnSpeed)

	drummers := map[string]int{"stitch": 626, "mimi": 627}
	fmt.Println("hello", drummers["stitch"])

	drummerItems := map[string]map[string]interface{}{
		"stitch": {"style": 626, "cymbal": "hhx"},
		"mimi":   {"style": 627, "cymbal": "ACustom"},
	}
	fmt.Println("hello world", drummerItems["stitch"]["cymbal"])

	drummerHandler := map[string]func(string) string{
		"stitch": func(style string) string {
			return style + " aka"
		},
		"mimi": func(style string) string {
			return style + " aka"
		},
	}

	for _, v := range []string{"stitch", "mimi"} {
		fmt.Println("return", drummerHandler[v]("metal"))
	}
}

func myHandler() {
	time.Sleep(time.Duration(2) * time.Second)
	music()
}

func HandleUuid() {
	uuid, _ := uuid.NewV4()
	fmt.Println("uuid=", uuid)
}

func sub1(queue <-chan int) {
	for {
		fmt.Println("sub1 received", <-queue)
	}
}

func sub2(queue <-chan int) {
	for {
		fmt.Println("sub2 received", <-queue)
	}
}

func main() {
	router := gin.Default()

	taskQueue := make(chan int)
	go sub1(taskQueue)
	go sub2(taskQueue)

	router.GET("/framyExam/", func(c *gin.Context) {
		name := c.Param("a")
		name := c.Param("b")
		taskQueue <- rand.Intn(100)
		GagService()
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


func main() {
	router := gin.Default()

	taskQueue := make(chan int)
	go sub1(taskQueue)
	go sub2(taskQueue)

	router.GET("/", func(c *gin.Context) {
		name := c.Param("name")
		taskQueue <- rand.Intn(100)
		GagService()
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
