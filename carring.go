package main

import "fmt"

// Замыкание. Возвращаемая функция может использовать переменную в родительской функции.
func multiply(x int) func(y int) int {
	return func(y int) int {
		return x * y
	}
}

func main() {
	// Каррирование функции
	var mult10 = multiply(10)

	fmt.Println(mult10(5))  // 50
	fmt.Println(mult10(3))  // 30
	fmt.Println(mult10(10)) // 100
}
