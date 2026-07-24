package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	jobCh := make(chan int, 10) // buffer แค่ 10 — สายพาน ไม่ใช่โกดัง
	resultCh := make(chan int, 10)

	var wg sync.WaitGroup

	// 1) spawn คนงาน "ก่อน" — ยืนรอที่ท่อพร้อมหยิบ
	for range 3 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(jobCh, resultCh)
		}()
	}

	// 2) คนป้อนงานแยกร่างไปเลย — main ไม่ต้องค้างที่ท่อเต็ม
	go func() {
		for i := range 1000 {
			jobCh <- i + 1 // ท่อเต็ม = block แป๊บ รอคนงานดึง แล้วไหลต่อ (backpressure!)
		}
		close(jobCh) // ผู้ส่งปิดท่อของตัวเอง
	}()

	// 3) ยามเฝ้าประตู — คนงานครบสามค่อยปิดท่อผล
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// 4) main ทำหน้าที่เดียว: เก็บผลจนท่อปิด
	count := 0
	for result := range resultCh {
		fmt.Println(result)
		count++
	}

	fmt.Printf("done: %d results in %v\n", count, time.Since(start))
}

func worker(jobCh <-chan int, resultCh chan<- int) {
	for j := range jobCh {
		resultCh <- j * 2 // เอา time.Sleep ออกก่อน จะได้เห็น 1000 งานจบไว
	}
}
