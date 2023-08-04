package main

import (
	"fmt"
)

func main() {
	testRemoveSmilesFromString()
}

func testRemoveSmilesFromString() {
	type testCase struct {
		str      string
		expected string
	}

	testCases := []testCase{
		{str: "I work in google :-)))", expected: "I work in google "},
		{str: "Cool :-) and I failed assesment there:-((", expected: "Cool  and I failed assesment there"},
		{str: "lol:)", expected: "lol:)"},
		{str: "YEEEEE BOIIII!!!! :-))(())", expected: "YEEEEE BOIIII!!!! (())"},
		{str: "Cringe :-)))))))))))))))", expected: "Cringe "},
	}

	passedCount := 0
	overallCount := len(testCases)
	for _, testCase := range testCases {
		if res := removeSmilesFromString(testCase.str); res != testCase.expected {
			fmt.Printf("input: %s; expected: %s; result: %s; \n", testCase.str, testCase.expected, res)
		} else {
			passedCount++
		}
	}

	fmt.Printf("Tests done - passed: %d; overall: %d; \n", passedCount, overallCount)
}

// Конечный автомат для состояний смайлика
func removeSmilesFromString(s string) string {
	const (
		start = iota
		needNose
		needMouth
		moreSad
		moreHappy
	)
	states := make([]map[rune]int, 5)
	for i := 0; i < 5; i++ {
		states[i] = make(map[rune]int)
	}
	// Конечный автомат. Ключ - текущее состояние, значение - мапа.
	// Ключ мапы - текущий символ, значение - следующее доступное состояние автомата.
	states[start][':'] = needNose
	states[needNose]['-'] = needMouth
	states[needMouth]['('] = moreSad
	states[needMouth][')'] = moreHappy
	states[moreSad]['('] = moreSad
	states[moreHappy][')'] = moreHappy
	// Результирующая строка (используем как стек).
	var stringStack []rune
	// Текущее состояние.
	curState := start
	for _, ch := range s {
		// Пытаемся получить следующее состояние, если текущий символ - символ из смайлика.
		nextState, isExist := states[curState][ch]
		// Добавляем символ если:
		// уже есть один или два символа из смайлика: ':', '-' (нужно вставить, т.к. после них может и не быть скобки, тогда это будет не смайлик)
		// ЛИБО не нашли новое состояние по текущему символу, т.е. символ не из смайлика (обычный символ - нужно вставить).
		if isExist && nextState <= needMouth || !isExist {
			stringStack = append(stringStack, ch)
		}
		// Если до этого был символ '-' и нашлось новое состояние, т.е. текущий символ ')' или '(' -
		// убираем последние два символа: ":-" (т.е. нашли смайлик - удалим его).
		if isExist && curState == needMouth {
			stringStack = stringStack[:len(stringStack)-2]
		}
		// Переопределяем новое состояние.
		// Если новое состояние не нашлось - переопределимся на состояние start.
		curState = nextState
		// При состояниях moreSad и moreHappy мы просто игнорируем скобки.
	}

	return string(stringStack)
}
