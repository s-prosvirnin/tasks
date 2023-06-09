package leet_code

// Скользящее окно

// скользящее окно
func maxVowels(s string, k int) int {
	// максимальное количество гласных
	maxVowels := 0
	// текущее количество гласных
	curVowels := 0
	// левая граница окна (индекс элемента в строке) или левый бегунок
	leftP := 0
	// мапа с гласными
	vowels := make(map[byte]struct{})
	// это сделано для удобства заполнения мапы
	for _, v := range [...]string{"a", "e", "i", "o", "u"} {
		vowels[v[0]] = struct{}{}
	}
	// rightP - правая граница окна (индекс элемента в строке) или правый бегунок
	for rightP := 0; rightP < len(s); rightP++ {
		// если правый бегунок гласная увеличим счетчик
		if _, ok := vowels[s[rightP]]; ok {
			curVowels++
		}
		// окно вышло за пределы k - уменьшим окно
		if rightP-leftP > k-1 {
			// если левый бегунок гласная уменьшим счетчик
			if _, ok := vowels[s[leftP]]; ok {
				curVowels--
			}
			// передвинем левый бегунок
			leftP++
		}
		// максимальное количество гласных
		if curVowels > maxVowels {
			maxVowels = curVowels
		}
	}

	return maxVowels
}
