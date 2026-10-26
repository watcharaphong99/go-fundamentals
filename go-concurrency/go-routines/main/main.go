package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	fmt.Println("start main()")

	go printStart()
	go printA2F()

	time.Sleep(50 * time.Millisecond)

	fmt.Print("สิ้นสุด main() \n\n")
}

func printA2F() {
	fmt.Println("เริ่มต้น printA2F()")

	a := 'A'
	f := a + 5

	for i := a; i <= f; i++ {
		time.Sleep(800 * time.Microsecond)
		fmt.Println(string(i))
	}

	fmt.Println("สิ้ยสุด printA2F")
}

func printStart() {
	fmt.Println("เริ่มต้น printStart()")

	for i := 1; i <= 5; i++ {
		time.Sleep(900 * time.Microsecond)
		fmt.Println(strings.Repeat("*", i))
	}

	fmt.Print("ิส้นสุด printStat()")
}
