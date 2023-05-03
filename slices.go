package main

import "fmt"

func main() {
	// a1()
	a2()
	// testMake()
}

func a2() {
	var x = []string{"A", "B", "C"}
	for i, s := range x {
		print(i, s, ",")
		x[i+1] = "M"
		x = append(x, "Z")
		x[i+1] = "Z"
	}
	print(len(x))
}

func a1() {
	a := [...]int{0, 1, 2, 3}
	x := a[:1]
	y := a[2:]
	x = append(x, y...)
	x = append(x, y...)
	fmt.Println(a, x)
}

func testMake() {
	a := make([]int, 3, 3)
	fmt.Println(a)
	b := make([]int, 3)
	fmt.Println(b)
}
