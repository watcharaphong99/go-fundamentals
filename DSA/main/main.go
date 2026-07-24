// package main

// import "fmt"

// func main() {
// 	a := []int{1, 2, 3}
// 	b := append(a[:1], 10)
// 	fmt.Println(b) //1 10
// 	fmt.Println(a) //1 10 3
// }

package main

import (
	"fmt"
	"sort"
)

func main() {

	ages := map[string]int{
		"charlie": 34,
		"alice":   31,
		"bob":     25,
	}

	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Println(name, ages[name])
	}

	// ages := make(map[string]int)
	// ages["bob"] = 25
	// fmt.Println(ages == nil)
	// fmt.Println(ages["bob"])
	// fmt.Println(len(ages))
	// delete(ages, "bob")
	// ages = make(map[string]int)

	// fmt.Println(ages)

	// ages := map[string]int{"baby": 0}

	// age, ok := ages["baby"]
	// test, oktest := ages["test"]
	// fmt.Println(age, ok)
	// fmt.Println(test, oktest)

	// ages := make(map[string]int)
	// ages["bob"]++
	// ages["bob"]--

	// fmt.Println(ages["bob"])
	// fmt.Println(ages["alice"])
	// a := make(map[string]int)
	// a["alice"] = 31

	// b := map[string]int{"bob": 25}

	// fmt.Println(a)
	// fmt.Println(b)

	// ages := map[string]int{"bob": 25}
	// ages["bob"] = 26
	// ages["alice"] = 31
	// fmt.Println("Ages", ages)

	// delete(ages, "bob")
	// fmt.Println("Ages2", ages)
}

// func countDuplicates(s string) map[rune]int {
// 	frequency := make(map[rune]int)

// 	for _, ch := range s {
// 		frequency[ch]++
// 		fmt.Println("frequency", frequency)
// 	}

// 	duplicates := make(map[rune]int)

// 	for ch, count := range frequency {
// 		if count > 1 {
// 			duplicates[ch] = count
// 		}
// 	}

// 	return duplicates
// }

// func main() {
// 	result := countDuplicates("programming")
// 	fmt.Println("result", result)
// }
