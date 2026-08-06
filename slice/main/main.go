package main

import "fmt"

func main() {
	//Slicetype เขียนเป็น `[]T` โดยไมื่มืีขนาดใน type จึงใช่้กับ sequenceหลายความยาวได้้

	// printAll([]int{1, 2})
	// printAll([]int{1, 2, 3, 4})

	// months := [...]string{
	// 	1: "January", 2: "February", 3: "March",
	// 	4: "April", 5: "May", 6: "June",
	// 	7: "July", 8: "August", 9: "September",
	// 	10: "October", 11: "November", 12: "December",
	// }

	// q2 := months[4:7]
	// summer := months[6:9]

	// fmt.Println(q2)
	// fmt.Println(summer)
	//รวม index เร้ิ่มตั้น แตั่ไมื่รวม index ปลายทาง : [4:7] คู่ือ 4,5,6

	// months := [...]string{
	// 	1: "January", 2: "February", 3: "March",
	// 	4: "April", 5: "May", 6: "June",
	// 	7: "July", 8: "August",
	// }
	// q2 := months[4:7]
	// summer := months[6:9]

	// summer[0] = "JUNE!"
	// fmt.Println("q2:", q2)
	// fmt.Println("summer:", summer)

	// q2: [April May JUNE!]
	// summer: [JUNE! July August]
	//`q2` กับ `summer` ซ้้อนทับกันที่ June การแก้ผ่าน Slice หนึ่งจึงเห็นผ่านอีก Slice

	//`len` คู่ือจํานวน element ที่มองเห็น สื่วน `cap` คู่ือระยะจากจุดเร้ิ่ม Slice ถูึงท้าย underlyingarray
	// s := []int{10, 20, 30, 40, 50}
	// x := s[1:3]

	// fmt.Println(x)
	// fmt.Println("len:", len(x))
	// fmt.Println("cap:", cap(x))

	//Slice ขยายเกิน len เด้ิมได้้ หากยังไมื่เกิน cap
	// s := []int{10, 20, 30, 40, 50}
	// x := s[1:3]
	// y := x[:4]

	// fmt.Println("x", x)
	// fmt.Println("y", y)
	values := []int{1, 2, 3}
	chageFirst(values)
	fmt.Println(values)

}

func chageFirst(s []int) {
	s[0] = 9999
}

func printAll(values []int) {
	fmt.Println(values)
}
