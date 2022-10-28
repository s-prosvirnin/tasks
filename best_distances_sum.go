package main

import (
	"fmt"
)

func bestDistancesSum(maxDistance, maxCitiesCount int, citiesDistances []int) int {

	if len(citiesDistances) < maxCitiesCount {
		return -1
	}

	for _, distance := range citiesDistances {

	}

	return -1
}

func doBestDistancesSum(maxDistance, maxCitiesCount int, citiesDistances []int, expectedDistance int) {
	resultDistance := bestDistancesSum(maxDistance, maxCitiesCount, citiesDistances)
	if resultDistance != expectedDistance {
		fmt.Printf("TEST NOT PASSED. maxDistance: %d, maxCitiesCount: %d, citiesDistances: %v, expectedDistance: %d, resultDistance: %d\n", maxDistance, maxCitiesCount, citiesDistances, expectedDistance, resultDistance)

		return
	}

	fmt.Printf("test passed. maxDistance: %d, maxCitiesCount: %d, citiesDistances: %v, expectedDistance: %d, resultDistance: %d\n", maxDistance, maxCitiesCount, citiesDistances, expectedDistance, resultDistance)
}

func main() {
	doBestDistancesSum(
		163,
		3,
		[]int{50, 55, 56, 57, 58},
		163,
	)
	doBestDistancesSum(163, 3, []int{50}, -1)
	citiesDistances := []int{91, 74, 73, 85, 73, 81, 87}
	doBestDistancesSum(230, 3, citiesDistances, 228)
	doBestDistancesSum(331, 2, citiesDistances, 178)
	doBestDistancesSum(331, 4, citiesDistances, 331)
	doBestDistancesSum(331, 5, citiesDistances, -1)
}
