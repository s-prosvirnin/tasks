package main

import "fmt"

type myType struct{}

func isNotNil(arg interface{}) {
	if arg != nil {
		fmt.Println("second not nil")
	}
}

func main() {
	var test *myType
	if test != nil {
		fmt.Println("first not nil")
	}
	isNotNil(test)
}
