package main

import (
	"fmt"
	"time"
)

func main() {
	workerCout := 10

	// 1. Use channel to receive data from go routine
	//    Send data from go routine back to main

	responseChanel := make(chan string)

	for i := 0; i < workerCout; i++ {
		workerID := fmt.Sprintf("worker-%d", i)
		go worker1(workerID, responseChanel)
	}

	for i := 0; i < workerCout; i++ {
		res := <-responseChanel
		println(res)
	}

	close(responseChanel)
	println("All response returned")

	// 2. Use channel to signal go module to exit
	//    Send data from outside go model into module

	exitChanel := make(chan bool)
	for i := 0; i < workerCout; i++ {
		workerID := fmt.Sprintf("worker-%d", i)
		go worker3(workerID, exitChanel)
	}

	time.Sleep(10 * time.Second)
	for i := 0; i < workerCout; i++ {
		exitChanel <- true
	}
	close(exitChanel)
	time.Sleep(2 * time.Second)
	println("Main is exited")
}

func worker3(workerID string, exitChanel chan bool) {
	i := 0
	for true {
		i++

		select {
		case <-exitChanel:
			println(fmt.Sprintf("workerID=%s has exited", workerID))
			return
		default:
			println(fmt.Sprintf("Worker=%s, Counter=%d", workerID, i))
			time.Sleep(1 * time.Second)
		}

	}
}

func worker1(workerID string, responseChanel chan string) {
	time.Sleep(1 * time.Second)
	responseChanel <- (workerID + " Response")
}
