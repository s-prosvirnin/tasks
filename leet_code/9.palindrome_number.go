package leet_code

import "strconv"

//https://leetcode.com/problems/palindrome-number/description/

func isPalindrome(x int) bool {
	numsStr := strconv.Itoa(x)
	numsLen := len(numsStr)
	for pos, num := range numsStr {
		if num != rune(numsStr[numsLen-pos-1]) {
			return false
		}
	}

	return true
}
