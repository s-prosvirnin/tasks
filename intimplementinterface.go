package main

import "fmt"

func main() {
	c1 := (*counter)(new(int))
	needCallable1(c1)
	needCallable1(c1)
	needCallable1(c1)
	fmt.Println(*c1)
}

type callable1 interface {
	call()
}

type counter int

func (c *counter) call() {
	*c++
}

func needCallable1(c callable1) {
	c.call()
}
