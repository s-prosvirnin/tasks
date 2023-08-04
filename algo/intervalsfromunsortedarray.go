package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	testIntervalsFromUnsortedArray()
}

func testIntervalsFromUnsortedArray() {
	type testCase struct {
		nums     []int
		expected string
	}

	testCases := []testCase{
		{nums: []int{1, 4, 5, 2, 3, 9, 8, 11, 0}, expected: "0-5,8-9,11"},
		{nums: []int{1, 4, 3, 2}, expected: "1-4"},
		{nums: []int{1, 4}, expected: "1-4"},
	}

	passedCount := 0
	overallCount := len(testCases)
	for _, testCase := range testCases {
		if res := intervalsFromUnsortedArray(testCase.nums); res != testCase.expected {
			fmt.Printf("Expected: %s; result: %s; \n", testCase.expected, res)
		} else {
			passedCount++
		}
	}

	fmt.Printf("Tests done - passed: %d; overall: %d; \n", passedCount, overallCount)
}

// TODO: доработать решение, чтобы результирующие интервалы были отсортированы
func intervalsFromUnsortedArray(nums []int) string {
	// Мапа с числами, которые есть в массиве для быстрого поиска. Ключ - число, значение не важно.
	numsMap := make(map[int]struct{})
	// Результирующая строка с интервалами.
	res := strings.Builder{}
	// Заполняем мапу.
	for _, num := range nums {
		numsMap[num] = struct{}{}
	}
	// Бежим по мапе. startNum - потенциальное начало интервала.
	for startNum := range numsMap {
		// Если слева от числа не заполнено (т.е. число слева пропущено), то мы нашли интервал.
		if _, ok := numsMap[startNum-1]; !ok {
			// Потенциальный конец интервала.
			endNum := startNum
			// Бежим последовательно вправо от startNum, все последовательные числа входят в текущий интервал.
			// Если число пропущено, то завершаем цикл. Последнее число это конец интервала.
			for _, ok := numsMap[endNum+1]; ok; _, ok = numsMap[endNum+1] {
				endNum++
			}
			// Ставим запятую к предыдущему интервалу.
			if res.Len() > 0 {
				res.WriteByte(',')
			}
			// Заполняем начало интервала.
			res.WriteString(strconv.Itoa(startNum))
			// Если конец интервала больше чем начало - заполняем конец интервала.
			// Если интервал состоит из одного числа, то конец не заполняем.
			if endNum != startNum {
				res.WriteByte('-')
				res.WriteString(strconv.Itoa(endNum))
			}
		}
	}
	// Если нам нужны отсортированные интервалы, то нужно писать интервалы в массив,
	// а здесь отсортировать результирующий массив и превратить в строку.

	return res.String()
}
