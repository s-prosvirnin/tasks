package main

import "fmt"

func main() {
	counter1 := closureCounter()
	fmt.Println(counter1(1))
	fmt.Println(counter1(-2))
	fmt.Println(counter1(3))

	counter2 := closureCounter()
	fmt.Println(counter2(3))
}

// Замыкание. Возвращаемая функция может использовать переменную в родительской функции.
func closureCounter() func(incr int) int {
	counter := 0

	return func(incr int) int {
		counter += incr
		return counter
	}
}
