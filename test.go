package main

import (
	"fmt"
)

type X struct {
	V int
}

func (x X) S() {
	fmt.Println(x.V)
}

// что выведет данный код
func main() {
	x := X{123}
	defer x.S()
	x.V = 456
}
