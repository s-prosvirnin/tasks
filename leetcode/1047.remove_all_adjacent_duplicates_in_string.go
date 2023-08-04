package leetcode

/**
You are given a string s consisting of lowercase English letters. A duplicate removal consists of choosing two adjacent and equal letters and removing them.
We repeatedly make duplicate removals on s until we no longer can.
Return the final string after all such duplicate removals have been made. It can be proven that the answer is unique.

Input: s = "abbaca"
Output: "ca"
Explanation:
For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move.  The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".

Input: s = "azxxzy"
Output: "ay"
*/

func removeDuplicates(s string) string {
	// примитивная реализация стека со слайсом и бегунком. Вместо бегунка можно просто обрезать слайс.
	var stack []rune
	stackPos := -1
	for _, r := range s {
		// делаем фиктивный pop элемента из стека (передвигаем бегунок назад) если соседние буквы совпадают
		if stackPos >= 0 && r == stack[stackPos] {
			stackPos--
			continue
		}
		// буквы не совпадают - смещаем бегунок вперед и добавляем текущую букву в стек
		stackPos++
		// если бегунок уперся в длину слайса, то аппендим
		if stackPos == len(stack) {
			stack = append(stack, r)
			continue
		}
		// иначе просто заменяем элемент под бегунком
		stack[stackPos] = r
	}

	return string(stack[:stackPos+1])
}
