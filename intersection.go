package main

import "fmt"

//На вход подаются два неупорядоченных слайса любой длины. Надо написать функцию, которая возвращает их пересечение

func intersection(array1, array2 []int) []int {
	var valuesCountersMap = make(map[int]int, len(array1)+len(array2))
	var intersectionCollection []int
	for _, value := range array1 {
		if _, ok := valuesCountersMap[value]; ok {
			valuesCountersMap[value]++
		} else {
			valuesCountersMap[value] = 1
		}
	}
	for _, value := range array2 {
		if counter, ok := valuesCountersMap[value]; ok && counter > 0 {
			intersectionCollection = append(intersectionCollection, value)
			valuesCountersMap[value]--
		}
	}

	return intersectionCollection
}

func intersectionTest() {
	a := []int{23, 3, 1, 2}
	b := []int{6, 2, 4, 23}
	// [2, 23]
	fmt.Printf("%v\n", intersection(a, b))
	a = []int{1, 1, 1}
	b = []int{1, 1, 1, 1}
	// [1, 1, 1]
	fmt.Printf("%v\n", intersection(a, b))
}

func main() {
	intersectionTest()
}
