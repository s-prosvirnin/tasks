package leetcode

/*
*
https://leetcode.com/problems/fibonacci-number/
509. Fibonacci Number

The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,
F(0) = 0, F(1) = 1
F(n) = F(n - 1) + F(n - 2), for n > 1.
Given n, calculate F(n).
*/
func fib1(n int) int {
	prevElements := [2]int{1, 1}
	for i := 3; i <= n; i++ {
		prevElements[i%2] = prevElements[0] + prevElements[1]
	}

	return prevElements[n%2]
}

func fib2(n int) int {
	if n == 0 {
		return 0
	}
	n1 := 0
	n2 := 1
	for i := 2; i < n; i++ {
		n2Old := n2
		n2 = n1 + n2
		n1 = n2Old
	}

	return n2 + n1
}
