package main

import (
	"fmt"
	"math"
)

// โครงสร้างข้อมูลของแต่ละช่วงภาษี
type TaxBracket struct {
	Upper float64 // เพดานของช่วงรายได้
	Rate  float64 // อัตราภาษี
}

// ตารางภาษี
var brackets = []TaxBracket{
	{Upper: 80000, Rate: 0.00},
	{Upper: 100000, Rate: 0.05},
	{Upper: 500000, Rate: 0.10},
	{Upper: 1000000, Rate: 0.20},
	{Upper: math.MaxFloat64, Rate: 0.30},
}

// ฟังก์ชันคำนวณภาษี
func CalculateTax(income float64) float64 {

	if income < 0 {
		panic("เงินได้ต้องไม่ติดลบ")
	}

	tax := 0.0
	lower := 0.0

	// วนลูปทีละช่วงภาษี
	for _, bracket := range brackets {

		// ถ้ารายได้ไม่ถึงช่วงนี้แล้ว ให้หยุด
		if income <= lower {
			break
		}

		// คำนวณจำนวนเงินที่อยู่ในช่วงนี้
		taxable := math.Min(income, bracket.Upper) - lower

		// คำนวณภาษีของช่วงนี้
		taxInBracket := taxable * bracket.Rate

		// แสดงรายละเอียด
		fmt.Println("--------------------------------")
		fmt.Printf("ช่วงรายได้ : %.0f - %.0f\n", lower, bracket.Upper)
		fmt.Printf("อัตราภาษี : %.0f%%\n", bracket.Rate*100)
		fmt.Printf("เงินที่เสียภาษี : %.0f\n", taxable)
		fmt.Printf("ภาษีช่วงนี้ : %.0f\n", taxInBracket)

		// รวมภาษี
		tax += taxInBracket

		// เลื่อนไปช่วงถัดไป
		lower = bracket.Upper
	}

	return tax
}

func main() {

	income := 600000.0

	fmt.Printf("เงินได้สุทธิ : %.0f บาท\n\n", income)

	tax := CalculateTax(income)

	fmt.Println("--------------------------------")
	fmt.Printf("ภาษีทั้งหมด : %.0f บาท\n", tax)
}
