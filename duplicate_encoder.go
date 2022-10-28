package main

import (
	"fmt"
	"strings"
)

func DuplicateEncode(word string) string {
	word = strings.ToUpper(word)
	var symbolsRepeatMap = make(map[int32]int, len(word))
	for _, symbol := range word {
		symbolsRepeatMap[symbol]++
	}

	var duplicateEncodeString strings.Builder
	for _, symbol := range word {
		if symbolsRepeatMap[symbol] > 1 {
			duplicateEncodeString.WriteString(")")
		} else {
			duplicateEncodeString.WriteString("(")
		}
	}

	return duplicateEncodeString.String()
}

func doTestDuplicateEncode(word string, expected string) {
	if result := DuplicateEncode(word); result != expected {
		fmt.Printf("TEST NOT PASSED. word: %s, expected: %s, result: %s\n", word, expected, result)
	} else {
		fmt.Printf("test passed. word: %s, expected: %s, result: %s\n", word, expected, result)
	}
}

func main() {
	doTestDuplicateEncode("din", "(((")
	doTestDuplicateEncode("recede", "()()()")
	doTestDuplicateEncode("(( @", "))((")
	doTestDuplicateEncode("Success", ")())())")
	doTestDuplicateEncode("", "")
}
