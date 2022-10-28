package main

import (
	"fmt"
	"math"
	"strconv"
)

func digitPows(number int, firstDigitPow int) int {

	//(a ^ p + b ^ (p+1) + c ^(p+2) + d ^ (p+3) + ...)
	sumResult := 0
	iteratableNumber := number
	numberDigitsCount := len(strconv.Itoa(number)) - 1
	fmt.Printf("number: %d, firstDigitPow: %d\n", number, firstDigitPow)
	for ; numberDigitsCount >= 0; numberDigitsCount-- {
		iteratableNumber = number / int(math.Pow(float64(10), float64(numberDigitsCount)))
		//fmt.Printf("iteratableNumber: %d, number10: %d\n", iteratableNumber, iteratableNumber % 10)
		sumResult += int(math.Pow(float64(iteratableNumber%10), float64(firstDigitPow)))
		firstDigitPow++
	}

	multiplier := float64(sumResult) / float64(number)

	if multiplier >= 1 && multiplier == float64(int64(multiplier)) {
		return int(multiplier)
	}

	return -1
}

func doTestDigitPows(digit int, firstDigitPow int, expectedDigitMultiplier int) {
	digitMultiplier := digitPows(digit, firstDigitPow)
	if digitMultiplier != expectedDigitMultiplier {
		fmt.Printf("TEST NOT PASSED. digit: %d, firstDigitPow: %d, expectedDigitMultiplier: %d, resultDigitMultiplier: %d\n", digit, firstDigitPow, expectedDigitMultiplier, digitMultiplier)

		return
	}

	fmt.Printf("test passed. digit: %d, firstDigitPow: %d, expectedDigitMultiplier: %d, resultDigitMultiplier: %d\n", digit, firstDigitPow, expectedDigitMultiplier, digitMultiplier)
}

func main() {
	doTestDigitPows(89, 1, 1)
	doTestDigitPows(92, 1, -1)
	doTestDigitPows(695, 2, 2)
	doTestDigitPows(46288, 3, 51)
	doTestDigitPows(46288, 5, -1)
}
