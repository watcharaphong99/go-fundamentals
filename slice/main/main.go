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
	// values := []int{1, 2, 3}
	// chageFirst(values)
	// fmt.Println(values)

	//make สร้าง underlying array และคือ slice ที่มี len/cap ตามกำหนด
	// a := make([]int, 3)
	// b := make([]int, 3, 6)

	// fmt.Println(a, len(a), cap(a))
	// fmt.Println(b, len(b), cap(b))

	//===========================
	// x := make([]int, 3, 5)
	// copy(x, []int{1, 2, 3})

	// y := append(x, 100)
	// y[0] = 999

	// fmt.Println("x:", x)
	// fmt.Println("y:", y)

	//output x: [999 2 3] y: [999 2 3 100]
	//ห้ามเดาว่า append จะ share หร้ือแยก ให้คู่ิดเสมอว่า append อาจคู่ืน Slice ใหมื่

	/////////////////////////////////////////////////
	// สำคัญมาก function ถ้ามีการใช้ slice ต้อง return เสมอ
	//append ไม่ได้แก้ slice เดิมโดยตรง แต่ คืน slice ใหม่ — ถ้าไม่รับค่ากลับ caller จะไม่เห็นผล

	// values := []int{1, 2, 3}

	// addWrong(values, 4)
	// fmt.Println("wrong:", values)

	// values = addRight(values, 4)
	// fmt.Println("right", values)

	//// 1) return slice ใหม่ (นิยมที่สุด)
	// func addRight(s []int, v int) []int {
	//     return append(s, v)
	// }

	// // 2) รับ pointer to slice
	// func addWithPointer(s *[]int, v int) {
	//     *s = append(*s, v)
	// }

	// // 3) pre-allocate แล้ว return
	// func build() []int {
	//     s := make([]int, 0, 100)
	//     // ...
	//     return s
	// }
	/////////////////////////////////////////////////

	//ทั้งคูู่่ len=0 แตั่ nilSlice ไมื่มืี underlyingarray สื่วน emptySlice เป็นคู่่าที่ไมื่ nil
	// var a []int
	// b := []int{}

	// fmt.Println(len(a), a == nil)
	// fmt.Println(len(b), b == nil)

	// b = append(b, 1)
	// fmt.Println(b)

	///////////////////////
	// Slice เหมาะกับ algorithm ที่แก้ underlyingarray โดยไมื่ allocateใหม่

	// values := []int{1, 2, 3, 4, 5}
	// reverse(values)
	// fmt.Println(values)

	// values := []int{1, 0, 2, 0, 3, 4, 5, 0}
	// result := noneZero(values)
	// fmt.Println(result)

	a := []int{5, 6, 7, 8, 9}
	b := []int{5, 6, 7, 8, 9}
	fmt.Println(removeOrdered(a, 2))
	fmt.Println(removeFast(b, 2))

}

func removeOrdered(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func removeFast(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func noneZero(values []int) []int {
	out := values[:0]

	for _, v := range values {
		if v != 0 {
			out = append(out, v)
		}
	}
	return out
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func addWrong(s []int, value int) {
	s = append(s, value)
}

func addRight(s []int, value int) []int {
	return append(s, value)
}

func chageFirst(s []int) {
	s[0] = 9999
}

func printAll(values []int) {
	fmt.Println(values)
}
