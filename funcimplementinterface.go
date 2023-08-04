package main

import "fmt"

func main() {
	var c1 caller1 = func() {
		fmt.Println("call1")
	}
	needCallable(c1)
	var c2 caller2
	needCallable(c2)
}

type callable2 interface {
	call()
}

type caller1 func()

func (c caller1) call() {
	c()
}

type caller2 func()

func (c caller2) call() {
	fmt.Println("call2")
}

func needCallable(c callable2) {
	c.call()
}
