package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	testArray1DiffArray2()
}

func testArray1DiffArray2() {
	type testCase struct {
		nums1    []int
		nums2    []int
		expected []int
	}

	testCases := []testCase{
		{nums1: []int{1, 2, 4}, nums2: []int{1, 2, 3, 4, 5}, expected: []int{}},
		{nums1: []int{1, 2, 3, 5, 6}, nums2: []int{1, 2, 4}, expected: []int{3, 5, 6}},
		{nums1: []int{1, 2}, nums2: []int{}, expected: []int{1, 2}},
		{nums1: []int{}, nums2: []int{1, 2}, expected: []int{}},
		{nums1: []int{}, nums2: []int{}, expected: []int{}},
		{nums1: []int{0, 3, 3, 3}, nums2: []int{1, 3}, expected: []int{0}},
		{nums1: []int{1, 2, 2, 3}, nums2: []int{1, 2, 4}, expected: []int{3}},
		{nums1: []int{10, 20, 30}, nums2: []int{19, 20, 21}, expected: []int{10, 30}},
		{nums1: []int{1, 2, 3, 4}, nums2: []int{19, 20, 21, 22}, expected: []int{1, 2, 3, 4}},
		{nums1: []int{19, 20, 21, 22}, nums2: []int{1, 2, 3, 4}, expected: []int{19, 20, 21, 22}},
	}

	passedCount := 0
	overallCount := len(testCases)
	for _, testCase := range testCases {
		if res := array1DiffArray2(testCase.nums1, testCase.nums2); !slices.Equal(testCase.expected, res) {
			fmt.Printf("Expected: %d; result: %d; \n", testCase.expected, res)
		} else {
			passedCount++
		}
	}

	fmt.Printf("Tests done - passed: %d; overall: %d; \n", passedCount, overallCount)
}

// Важные кейсы:
// элементы второго массива закончились, а первого нет
// элементы первого массива закончились, а второго нет
// nums1 = [1,2,2,2,3], nums2 = [1,2]
// nums1 = [0,3,3], nums2 = [1,3]
func array1DiffArray2(nums1 []int, nums2 []int) []int {
	// Результирующий массив.
	var res []int
	// Бегунки для первого и второго массивов.
	idx1, idx2 := 0, 0
	// Пока не вышли за рамки одного из массивов.
	// Включаем кейс: элементы первого массива закончились, а второго нет.
	for idx1 < len(nums1) && idx2 < len(nums2) {
		// Если элемент второго массива меньше чем элемент первого -
		// сдвигаем бегунок второго массива и выходим из итерации.
		if nums2[idx2] < nums1[idx1] {
			idx2++
			continue
		}
		// Если элемент первого массива меньше чем элемент второго -
		// добавляем элемент первого массива (мы нашли искомый элемент).
		if nums1[idx1] < nums2[idx2] {
			res = append(res, nums1[idx1])
		}
		// Сюда пришли когда элемент первого массива не больше (<=) элемента второго массива.
		// Смещаем бегунок первого массива.
		// Если элементы равны, то смещать второй бегунок не нужно, иначе не покроем кейс:
		// nums1 = [0,3,3], nums2 = [1,3].
		idx1++
	}
	// Элементы второго массива закончились, а первого нет.
	// В первом цикле обработать этот кейс более сложно, чем отдельно.
	for ; idx1 < len(nums1); idx1++ {
		// Если второй массив пуст или элемент из первого массива больше чем второй -
		// добавляем элемент в результирующий массив.
		// Покрываем кейс: nums1 = [0,3,3,3], nums2 = [1,3].
		if len(nums2) == 0 || nums1[idx1] > nums2[len(nums2)-1] {
			res = append(res, nums1[idx1])
		}
	}

	return res
}
