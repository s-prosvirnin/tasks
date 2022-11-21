package main

import "fmt"

func main() {
	//rangeScope()
	range1()
}

func rangeScope() {
	a := [2]int{1, 2}

	fmt.Print("Pass 1: ")
	// копируем массив в скоуп rang-а (по сути range своего рода функция)
	for _, v := range a {
		// изменяем только внешний массив, поэтому внутри цикла элемент не меняется
		a[1] = 10
		fmt.Printf("%d, ", v)
	}
	// сбросим измененное значение
	a[1] = 2

	fmt.Print("  Pass 2: ")
	// копируем указатель на массив в скоуп rang-а
	b := &a
	for _, v := range b {
		// внешний массив передан по ссылке, поэтому элемент изменится
		b[1] = 10
		fmt.Printf("%d, ", v)
	}
}

func range1() {
	x := []int{7, 8, 9}
	y := [3]*int{}
	for i, v := range x {
		defer print(v)

		y[i] = &i
	}

	print(*y[0], *y[1], *y[2], " ")
}
