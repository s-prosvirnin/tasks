package leet_code

/*
	Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

	Input: n = 3
	Output: ["((()))","(()())","(())()","()(())","()()()"]

	Input: n = 1
	Output: ["()"]
*/

func generateParenthesis(n int) []string {
	return a("", n)
}

func a(s string, level int) []string {
	if level == 0 {
		return []string{s}
	}
	var ss []string
	level--
	for ; level >= 0; level-- {
		for _, sss := range a("()"+s, level) {
			ss = append(ss, sss)
		}
		for _, sss := range a("("+s, level) {
			ss = append(ss, sss+")")
		}
	}

	return ss
}
