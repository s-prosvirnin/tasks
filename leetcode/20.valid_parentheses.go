package leetcode

// https://leetcode.com/problems/valid-parentheses

func isValid(s string) bool {
	stack := make([]string, len(s))
	stackPointer := -1
	res := true
	for _, par := range s {
		if string(par) == ")" {
			if stackPointer < 0 || stack[stackPointer] != "(" {
				res = false
				break
			}
			stackPointer--
			continue
		}
		if string(par) == "}" {
			if stackPointer < 0 || stack[stackPointer] != "{" {
				res = false
				break
			}
			stackPointer--
			continue
		}
		if string(par) == "]" {
			if stackPointer < 0 || stack[stackPointer] != "[" {
				res = false
				break
			}
			stackPointer--
			continue
		}
		stackPointer++
		stack[stackPointer] = string(par)
	}

	if stackPointer >= 0 {
		res = false
	}

	return res
}
