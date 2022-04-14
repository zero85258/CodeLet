package main

import "fmt"
import "time"

func main() {
	fmt.Println("Hello world!")
	goal := make(chan string)

	fmt.Println(goal)
	c1 := make(chan string, 2)
	c2 := make(chan string, 2)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	// consumer
	func() {
		for i := 0; i < 2; i++ {
			select {
			case msg1 := <-c1:
				fmt.Println("received", msg1)
			case msg2 := <-c2:
				fmt.Println("received", msg2)
			}
		}
	}()
}
