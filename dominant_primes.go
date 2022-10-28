package main

import (
	"fmt"
	"math"
	"math/big"
)

/**
The prime number sequence starts with: 2,3,5,7,11,13,17,19.... Notice that 2 is in position one.
3 occupies position two, which is a prime-numbered position. Similarly, 5, 11 and 17 also occupy prime-numbered positions. We shall call primes such as 3,5,11,17 dominant primes because they occupy prime-numbered positions in the prime number sequence. Let's call this listA.
As you can see from listA, for the prime range range(0,10), there are only two dominant primes (3 and 5) and the sum of these primes is: 3 + 5 = 8.
Similarly, as shown in listA, in the range (6,20), the dominant primes in this range are 11 and 17, with a sum of 28.
Given a range (a,b), what is the sum of dominant primes within that range? Note that a <= range <= b and b will not exceed 500000.
*/

//2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199

func dominantPrimesRangeSum2(startNumber, endNumber int) int {
	primeNumbersAsKeysArray := make([]int, endNumber+1)
	primeNumbersArrayKey := 0
	for number := 2; number <= endNumber; number++ {
		//if big.NewInt(int64(number)).ProbablyPrime(0) {
		if isPrime(number) {
			primeNumbersArrayKey++
			primeNumbersAsKeysArray[primeNumbersArrayKey] = number
		}
	}

	dominantPrimesRangeSum := 0
	for _, primaryNumber := range primeNumbersAsKeysArray {
		if primaryNumber > 0 && primeNumbersAsKeysArray[primaryNumber] > startNumber {
			dominantPrimesRangeSum += primeNumbersAsKeysArray[primaryNumber]
		}
	}

	return dominantPrimesRangeSum
}

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n < 2 || n&1 == 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeNumbersGenerator2(endNumber int) []int {
	primeNumbersAsKeysArray := make([]int, endNumber+1)
	primeNumbersArrayKey := 0
	for number := 2; number <= endNumber; number++ {
		if big.NewInt(int64(number)).ProbablyPrime(0) {
			primeNumbersArrayKey++
			primeNumbersAsKeysArray[primeNumbersArrayKey] = number
		}
	}

	primeNumbersArray := make([]int, primeNumbersArrayKey)
	primeNumbersArrayKey = -1
	for _, primaryNumber := range primeNumbersAsKeysArray {
		if primaryNumber > 0 {
			primeNumbersArrayKey++
			primeNumbersArray[primeNumbersArrayKey] = primaryNumber
		}
	}

	return primeNumbersArray
}

func dominantPrimesRangeSum(startNumber, endNumber int) int {
	primeNumbers := primeNumbersGenerator(endNumber)
	dominantPrimeNumbersSum := 0
	for _, number := range primeNumbers {
		if len(primeNumbers)-1 >= number {
			if primeNumbers[number-1] >= startNumber {
				dominantPrimeNumbersSum += primeNumbers[number-1]
			}
		}
	}

	return dominantPrimeNumbersSum
}

func primeNumbersGenerator(endNumber int) []int {
	primeNumbersAsKeysArray := make([]int, endNumber+1)

	divisor := 2
	primeNumbersCount := 0
	stopNumber := int(math.Sqrt(float64(endNumber)))
	for keyAsNumber, flag := range primeNumbersAsKeysArray {
		if keyAsNumber < 2 {
			primeNumbersAsKeysArray[keyAsNumber] = -1
			continue
		}
		if flag == -1 {
			continue
		} else {
			divisor = keyAsNumber
		}
		if keyAsNumber <= stopNumber {
			for number := divisor + 1; number <= endNumber; number++ {
				if number%divisor == 0 {
					primeNumbersAsKeysArray[number] = -1
				}
			}
		}
		if primeNumbersAsKeysArray[keyAsNumber] == 0 {
			primeNumbersCount++
		}
	}

	primeNumbersArray := make([]int, primeNumbersCount)
	primeNumbersArrayKey := -1
	for keyAsNumber, flag := range primeNumbersAsKeysArray {
		if flag == 0 {
			primeNumbersArrayKey++
			primeNumbersArray[primeNumbersArrayKey] = keyAsNumber
		}
	}

	return primeNumbersArray
}

func showPrimeNumbers(end int) {
	for _, primeNumber := range primeNumbersGenerator2(end) {
		fmt.Printf("%d, ", primeNumber)
	}
}

func doTestDominantPrimes(start int, end int, expectedSum int) {
	if dominantPrimesRangeSum := dominantPrimesRangeSum2(start, end); dominantPrimesRangeSum != expectedSum {
		fmt.Printf(
			"TEST NOT PASSED. %d-%d. expected: %d, result: %d\n",
			start,
			end,
			expectedSum,
			dominantPrimesRangeSum,
		)
	} else {
		fmt.Printf("test passed. %d-%d. expected: %d, result: %d\n", start, end, expectedSum, dominantPrimesRangeSum)
	}
}

func doTestDominantPrimes2(start int, end int) {
	dominantPrimesRangeSum := dominantPrimesRangeSum(start, end)
	dominantPrimesRangeSum2 := dominantPrimesRangeSum2(start, end)
	if dominantPrimesRangeSum != dominantPrimesRangeSum2 {
		fmt.Printf(
			"TEST NOT PASSED. %d-%d. expected: %d, result: %d\n",
			start,
			end,
			dominantPrimesRangeSum2,
			dominantPrimesRangeSum,
		)
	} else {
		fmt.Printf(
			"test passed. %d-%d. expected: %d, result: %d\n",
			start,
			end,
			dominantPrimesRangeSum2,
			dominantPrimesRangeSum,
		)
	}
}

func doTestPrimaryNumbersGenerator(endNumber int) {
	primeNumbers1 := primeNumbersGenerator(endNumber)
	primeNumbers2 := primeNumbersGenerator2(endNumber)
	if len(primeNumbers1) != len(primeNumbers2) {
		fmt.Printf("wrong length")
	}
	for key, primeNumber := range primeNumbers1 {
		if primeNumbers2[key] != primeNumber {
			fmt.Printf("diff")
		}
	}
}

func main() {
	//showPrimeNumbers(200)
	//doTestPrimaryNumbersGenerator(500000)
	//return

	//doTestDominantPrimes2(13, 129999)
	//return

	/**
	TODO: не работает на рандомных
	*/
	doTestDominantPrimes(-20, 10, 8)
	doTestDominantPrimes(0, 10, 8)
	doTestDominantPrimes(6, 20, 28)
	doTestDominantPrimes(2, 200, 1080)
	doTestDominantPrimes(2, 199, 1080)
	doTestDominantPrimes(1000, 100000, 52114889)
	doTestDominantPrimes(4000, 500000, 972664400)
	doTestDominantPrimes(4000, 500000, 972664400)
	doTestDominantPrimes(4000, 500000, 972664400)
	doTestDominantPrimes(4000, 500000, 972664400)
}
