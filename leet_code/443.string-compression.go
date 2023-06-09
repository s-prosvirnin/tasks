package leet_code

import "strconv"

// https://leetcode.com/problems/string-compression/

func compress1(chars []byte) int {
	currentChar := chars[0]
	char := chars[0]
	charsPointer := 0
	charCounter := 0
	charsCount := len(chars)
	for pos := 0; pos <= charsCount; pos++ {
		if pos == charsCount {
			char = chars[pos-1]
		} else {
			char = chars[pos]
			if char == currentChar {
				charCounter++
				continue
			}
		}
		chars[charsPointer] = currentChar
		if charCounter > 1 {
			for _, digit := range strconv.Itoa(charCounter) {
				charsPointer++
				chars[charsPointer] = byte(digit)
			}
		}
		currentChar = char
		charsPointer++
		charCounter = 1
	}

	return charsPointer
}

func compress2(chars []byte) int {
	letterCount := 0
	letterPos := 0
	totalCount := 0
	for i, c := range chars {
		if chars[letterPos] == c {
			letterCount++
			continue
		}
		totalCount++
		if letterCount > 1 {
			totalCount += len(strconv.Itoa(letterCount))
		}
		letterCount = 1
		letterPos = i
	}
	totalCount++
	if letterCount > 1 {
		totalCount += len(strconv.Itoa(letterCount))
	}

	return totalCount
}
