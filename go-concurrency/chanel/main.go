package main

import "fmt"

func main() {
	ch := make(chan string)
	ch1 := make(chan string)

	go func() {
		ch1 <- "Hello from Best"
	}()

	go func() {
		ch <- "Hello from sender!"
	}()

	msg := <-ch
	msg1 := <-ch1
	fmt.Println(msg)
	fmt.Println(msg1)
}
