package main

/** Решение через типизированные функции и передачу в них всей строки **/
func isBeautyF(str string) bool {
	if len(str) == 0 {
		return false
	}

	rules := []isBeautyFT{isSameF, isAcsF, isDescF}

	for _, r := range rules {
		if r(str) == true {
			return true
		}
	}

	return false
}

type isBeautyFT func(str string) bool

func isSameF(str string) bool {
	var sameChar rune
	for _, c := range str {
		if sameChar == 0 {
			sameChar = c
			continue
		}
		if sameChar != c {
			return false
		}
	}

	return true
}

func isAcsF(str string) bool {
	if len(str) >= 6 {
		return false
	}

	var prevChar rune
	for _, c := range str {
		if prevChar == 0 {
			prevChar = c
			continue
		}
		if prevChar != c-1 {
			return false
		}
		prevChar = c
	}

	return true
}

func isDescF(str string) bool {
	if len(str) >= 6 {
		return false
	}

	var prevChar rune
	for _, c := range str {
		if prevChar == 0 {
			prevChar = c
			continue
		}
		if prevChar != c+1 {
			return false
		}
		prevChar = c
	}

	return true
}

/** Решение через интерфейс и передачу в него элемента строки **/
func isBeautyI(str string) bool {
	if len(str) == 0 {
		return false
	}

	rules := []isBeautyIT{isSameChecker{}, isAscChecker{}, isDescChecker{}}
	results := make([]bool, len(rules))

	for _, c := range str {
		for i, r := range rules {
			results[i] = r.isBeauty(byte(c))
		}
	}

	for _, r := range results {
		if r == true {
			return true
		}
	}

	return false
}

type isBeautyIT interface {
	isBeauty(char byte) bool
}

type isSameChecker struct {
	sameChar byte
}

func (c isSameChecker) isBeauty(char byte) bool {
	if c.sameChar == 0 {
		c.sameChar = char
		return true
	}
	if c.sameChar != char {
		return false
	}

	return true
}

type isAscChecker struct {
	prevChar   byte
	charsCount int
}

func (c isAscChecker) isBeauty(char byte) bool {
	c.charsCount++
	if c.charsCount >= 6 {
		return false
	}

	if c.prevChar == 0 {
		c.prevChar = char
		return true
	}
	if c.prevChar != char-1 {
		return false
	}
	c.prevChar = char

	return true
}

type isDescChecker struct {
	prevChar   byte
	charsCount int
}

func (c isDescChecker) isBeauty(char byte) bool {
	c.charsCount++
	if c.charsCount >= 6 {
		return false
	}

	if c.prevChar == 0 {
		c.prevChar = char
		return true
	}
	if c.prevChar != char+1 {
		return false
	}
	c.prevChar = char

	return true
}
