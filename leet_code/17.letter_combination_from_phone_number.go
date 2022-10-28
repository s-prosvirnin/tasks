package leet_code

import "bytes"

/**
https://leetcode.com/problems/letter-combinations-of-a-phone-number/
17. Letter Combinations of a Phone Number

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.
A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Решение. Решил через последовательное определение комбинации. Более изящное решение через рекурсию.
*/
func letterCombinations(digits string) []string {
	m := map[int][]string{
		2: []string{"a", "b", "c"},
		3: []string{"d", "e", "f"},
		4: []string{"g", "h", "i"},
		5: []string{"j", "k", "l"},
		6: []string{"m", "n", "o"},
		7: []string{"p", "q", "r", "s"},
		8: []string{"t", "u", "v"},
		9: []string{"w", "x", "y", "z"},
	}
	lettersCount := 0
	maxLettersInDigit := 0
	for _, digit := range digits {
		if len(m[int(digit)-'0']) > maxLettersInDigit {
			maxLettersInDigit = len(m[int(digit)-'0'])
		}
		if lettersCount == 0 {
			lettersCount = len(m[int(digit)-'0'])
			continue
		}
		lettersCount *= len(m[int(digit)-'0'])
	}
	res := make([]string, lettersCount)

	var letterSkip, letterFrequency, leftPos, rightPos, lettersInDigitsAccum int
	for digitPos, digit := range digits {
		lettersInDigit := len(m[int(digit)-'0'])
		if digitPos == 0 {
			lettersInDigitsAccum = lettersInDigit
			letterSkip = lettersInDigit - 1
			letterFrequency = 1
			leftPos = 0
			rightPos = letterFrequency - 1
		} else {
			letterFrequency = lettersInDigitsAccum
			letterSkip = letterFrequency * (lettersInDigit - 1)
			leftPos = 0
			rightPos = letterFrequency - 1
			lettersInDigitsAccum *= lettersInDigit
		}
		for letterPos, letter := range m[int(digit)-'0'] {
			for resPos := 0; resPos < lettersCount; resPos++ {
				if resPos < leftPos {
					continue
				}
				if resPos > rightPos {
					leftPos = rightPos + letterSkip + 1
					rightPos = (leftPos + letterFrequency) - 1
					continue
				}
				var buffer bytes.Buffer
				buffer.WriteString(res[resPos])
				buffer.WriteString(letter)
				res[resPos] = buffer.String()
			}
			leftPos = (letterFrequency * (letterPos + 1))
			rightPos = (leftPos + letterFrequency) - 1
		}
	}

	return res
}
