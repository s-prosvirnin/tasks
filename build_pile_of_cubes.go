package main

import (
	"fmt"
	"math"
)

func getCubesLevels(cubesCount int) int {

	var cubesLevels = -1
	var currentCubesCount = 0
	const MaxInt = int((^uint(0)) >> 1)
	for level := 1; currentCubesCount <= cubesCount; level++ {
		if currentCubesCount == cubesCount {
			break
		}
		currentCubesCount += int(math.Pow(float64(level), float64(3)))
		if currentCubesCount > cubesCount {
			cubesLevels = -1
			break
		}
		cubesLevels = level
	}

	return cubesLevels
}

func doTestCubesLevels(cubesCount int, expectedCubesLevels int) {
	cubesLevelsResult := getCubesLevels(cubesCount)
	if cubesLevelsResult != expectedCubesLevels {
		fmt.Printf("TEST NOT PASSED. cubesCount: %d, expectedCubesLevels: %d, cubesLevelsResult: %d\n", cubesCount, expectedCubesLevels, cubesLevelsResult)

		return
	}

	fmt.Printf("test passed. cubesCount: %d, expectedCubesLevels: %d, cubesLevelsResult: %d\n", cubesCount, expectedCubesLevels, cubesLevelsResult)
}

func main() {
	doTestCubesLevels(4183059834009, 2022)
	doTestCubesLevels(24723578342962, -1)
	doTestCubesLevels(91716553919377, -1)
	doTestCubesLevels(1071225, 45)
}
