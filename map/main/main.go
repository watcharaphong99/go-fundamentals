package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	fmt.Println("from==>", from)
	edges := graph[from]
	fmt.Println("edges==>", edges)

	if edges == nil {
		fmt.Println("edges1 in if==>", edges)
		edges = make(map[string]bool)
		fmt.Println("edges2 in if==>", edges)
		graph[from] = edges

		fmt.Println("graph", graph)
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

// var counts = make(map[string]int)

// func key(list []string) string {
// 	fmt.Printf("data==> %q", list)
// 	return fmt.Sprintf("%q", list)
// }

// func add(list []string) {
// 	fmt.Println("list==>", list)
// 	counts[key(list)]++
// }

// func equal(x, y map[string]int) bool {
// 	fmt.Println("x==>", x)
// 	fmt.Println("y==>", y)
// 	if len(x) != len(y) {
// 		return false
// 	}

// 	for k, value := range x {
// 		fmt.Println("key==>", k)
// 		fmt.Println("value==>", value)
// 		yv, ok := y[k]
// 		fmt.Println("ok", ok)
// 		fmt.Println("yv ", yv)
// 		if !ok || yv != value {
// 			return false
// 		}
// 	}
// 	return true
// }

func main() {
	//create map
	// a := make(map[string]int)
	// b := map[string]int{"best": 21}
	// a["bob"] = 23
	// fmt.Println(a)
	// fmt.Println(b)
	//==================================//
	//insert update delete key
	// ages := map[string]int{"bob": 3}

	// ages["bob"] = 2
	// ages["Best"] = 5
	// delete(ages, "bob")
	// fmt.Println(ages)
	//==================================//
	//Missing Key and Zero Value

	// ages := make(map[string]int)
	// ages["bob"]++
	// ages["bob"]++

	// fmt.Println(ages["bob"])
	// fmt.Println(ages["alice"])
	//==================================//
	//check value,ok in map
	// ages := map[string]int{"best": 21}
	// age, ok := ages["best"]

	// fmt.Println(age, ok)

	// ageNotFound, ok := ages["nobody"]
	// fmt.Println(ageNotFound, ok)
	//==================================//
	// Nil Map
	// var ages map[string]int

	// fmt.Println(ages == nil)
	// fmt.Println(ages["bob"])
	// fmt.Println(ages)
	// delete(ages, "bob")

	// ages = make(map[string]int)
	// ages["bob"] = 25
	// fmt.Println(ages)
	//==================================//
	//Range ไม่เรียงลำดับ
	//==================================//
	// ages := map[string]int{
	// 	"chalie": 34,
	// 	"alice":  31,
	// 	"bob":    25,
	// }

	// names := make([]string, 0, len(ages))

	// for name := range ages {
	// 	names = append(names, name)
	// }

	// sort.Strings(names)
	// for _, name := range names {
	// 	fmt.Println(name, ages[name])
	// }
	//==================================//

	// compaire map
	// a := map[string]int{"A": 0}
	// b := map[string]int{"A": 0}
	// c := map[string]int{"B": 0}

	// fmt.Println(maps.Equal(a, b))
	// fmt.Println(equal(a, c))

	//==================================//

	// map is set
	// seen := make(map[string]bool)

	// for _, value := range []string{"go", "rust", "go", "java"} {
	// 	if !seen[value] {
	// 		seen[value] = true
	// 		fmt.Println(value)
	// 	}

	// }

	//==================================//

	// count map

	// counts := make(map[string]int)

	// for _, word := range strings.Fields("go is simple and go is fase") {
	// 	counts[word]++
	// }

	// fmt.Println(counts["go"])
	// fmt.Println(counts["is"])

	//==================================//
	//slice not key

	// add([]string{"go", "map"})
	// // add([]string{"go", "map"})
	// fmt.Println(counts[key([]string{"go", "map"})])
	//==================================//
	//Nested Map is Graph
	addEdge("A", "B")
	addEdge("A", "C")

	fmt.Println(hasEdge("A", "B"))
	fmt.Println(hasEdge("A", "D"))
}
