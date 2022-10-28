package main

import "fmt"

func Snail(matrix [][]int) []int {
	matrixSize := len(matrix)

	if len(matrix[0]) == 0 {
		return []int{}
	}

	var resultArray = make([]int, matrixSize*matrixSize)
	resultArrayKey := -1
	for offset := 0; offset <= matrixSize-offset; offset++ {
		for key := offset; key <= matrixSize-offset-1; key++ {
			resultArrayKey++
			resultArray[resultArrayKey] = matrix[offset][key]
		}

		if offset == matrixSize-offset-1 {
			break
		}

		for key := offset + 1; key <= matrixSize-offset-2; key++ {
			resultArrayKey++
			resultArray[resultArrayKey] = matrix[key][matrixSize-offset-1]
		}

		for key := matrixSize - offset - 1; key >= offset; key-- {
			resultArrayKey++
			resultArray[resultArrayKey] = matrix[matrixSize-offset-1][key]
		}

		for key := matrixSize - offset - 2; key >= offset+1; key-- {
			resultArrayKey++
			resultArray[resultArrayKey] = matrix[key][offset]
		}
	}

	return resultArray
}

func snailTest() {
	snailMap1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Printf("%v\n", Snail(snailMap1))

	snailMap2 := [][]int{{1, 2, 3, 1}, {4, 5, 6, 4}, {7, 8, 9, 7}, {7, 8, 9, 7}}
	fmt.Printf("%v\n", Snail(snailMap2))

	snailMap3 := [][]int{{1, 2, 7, 4, 5}, {1, 2, 8, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 8, 4, 5}, {1, 2, 7, 4, 5}}
	fmt.Printf("%v\n", Snail(snailMap3))

	snailMap4 := [][]int{{1}}
	fmt.Printf("%v\n", Snail(snailMap4))

	snailMap5 := [][]int{{}}
	fmt.Printf("%v\n", Snail(snailMap5))
}

func main() {
	snailTest()
}
