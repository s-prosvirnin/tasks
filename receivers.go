package main

import "fmt"

func main() {
	s1 := s1{}
	s1.f1()
	s1.f2()
	s1.f3()
}

type s1 struct {
	a int
}

func (r s1) f1() {
	fmt.Println("f1. a=%d", r.a)
	r.a = 1
	fmt.Println("f1. a=%d", r.a)
}
func (r *s1) f2() {
	fmt.Println("f2. a=%d", r.a)
	r.a = 2
	fmt.Println("f2. a=%d", r.a)
}
func (s1) f3() {
	fmt.Println("f3")
}
