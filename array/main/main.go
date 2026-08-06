package main

import (
	"fmt"
)

func main() {
	// var a [3]int
	// fmt.Println(a)
	// fmt.Println("len", len(a))

	// a := [3]int{10, 20, 30}
	// for i, v := range a {
	// 	fmt.Printf("index=%d value=%d\n", i, v)
	// }

	// q := [...]int{1, 2, 3}
	// fmt.Printf("value=%v type=%T\n", q, q)

	// a := [3]int{1, 2, 3}
	// b := [4]int{1, 2, 3, 4}
	// a = b
	//cannot use b (variable of type [4]int) as [3]int value in assignment

	// a := [3]int{1, 2, 3}
	// b := a
	// b[0] = 999

	// fmt.Println("a:", a)
	// fmt.Println("b:", b)
	// a: [1 2 3]
	// b: [999 2 3]

	// a := [2]int{1, 2}
	// b := [...]int{1, 2}
	// c := [2]int{1, 3}

	// fmt.Println(a == b)
	// fmt.Println(a == c)

	// a := [3]int{1, 2, 3}
	// change(a)
	// fmt.Println(a)
	// fmt.Println("outside:", a)
	// inside: [999 2 3]
	// [1 2 3]
	// outside: [1 2 3]
	// coppy value ไม่เปลี่ยนต้นฉบับ ถ้า มี pointer เปลี่ยนต้นฉบับ

	// a := [3]int{1, 2, 3}
	// changeWithPointer(&a)
	// fmt.Println(a)

	// x := sha256.Sum256([]byte("x"))
	// X := sha256.Sum256([]byte("X"))

	// fmt.Printf("type =%T\n", x)
	// fmt.Println("equal:", x == X)
	//type=[32]uint8
	//equal: false

}

func changeWithPointer(a *[3]int) {
	a[0] = 999
	fmt.Println("inside:", a)
}

// func change(a [3]int) {
// 	a[0] = 999
// 	fmt.Println("inside:", a)
// }
