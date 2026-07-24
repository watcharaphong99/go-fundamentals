package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu      sync.Mutex
	balance int = 1000
	wg      sync.WaitGroup
)

func main() {

	// doneCh := make(chan bool, 3)

	// go UpdateBalance(doneCh, 100)
	// go UpdateBalance(doneCh, 500)
	// go UpdateBalance(doneCh, 400)

	// <-doneCh
	// <-doneCh
	// <-doneCh

	jobCh := make(chan int, 10)
	resultCh := make(chan int, 10)

	for i := range 10 {
		jobCh <- i + 1
	}

	close(jobCh)

	for range 3 {
		wg.Add(1)
		// go double(jobCh, resultCh)
		go func() {
			defer wg.Done() // คนงานรายงานตัวตอนเลิกงาน
			double(jobCh, resultCh)
		}()
	}

	go func() { // ← เงาเฝ้าประตูอีกหนึ่งร่าง
		wg.Wait()       // หลับรอจนคนงานครบ 3
		close(resultCh) // ครบแล้ว → ปิดท่อ
	}()

	for result := range resultCh { // range ได้เลย ไม่ต้องนับ 10 เอง!
		fmt.Println(result)
	}

	// for range 10 {
	// 	result := <-resultCh
	// 	fmt.Println(result)
	// }

}

// func UpdateBalance(doneCh chan<- bool, amount int) {

// 	mu.Lock()
// 	defer mu.Unlock()
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("update balance")

// 	balance -= amount
// 	doneCh <- true

// 	fmt.Println("Balance Update")

// }

func double(jobCh <-chan int, resultCh chan<- int) {
	for j := range jobCh {
		time.Sleep(1 * time.Second)
		resultCh <- j * 2
	}
}
