package leet_code

// https://leetcode.com/problems/score-of-parentheses/

func scoreOfParentheses(s string) int {
	stack := make([]int, len(s))
	stackPointer := 0
	//resTotal := 0
	//resTemp := 0
	for _, par := range s {
		// (
		if string(par) == "(" {
			stackPointer++
			continue
		}
		// )
		stackPointer--
		currentLevelVal := 0
		if stack[stackPointer+1] > 0 {
			currentLevelVal = stack[stackPointer+1] * 2
			stack[stackPointer+1] = 0
		} else {
			currentLevelVal = 1
		}
		stack[stackPointer] += currentLevelVal
	}

	return stack[stackPointer]
}
