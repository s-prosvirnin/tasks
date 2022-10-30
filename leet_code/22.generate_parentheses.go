package leet_code

/*
	Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

	Input: n = 3
	Output: ["((()))","(()())","(())()","()(())","()()()"]

	Input: n = 1
	Output: ["()"]
*/

func generateParenthesis(n int) []string {
	return a("", n, 0, 0, 0)
}

func a(s string, n int, brackets int, opened int, toClose int) []string {
	// достигаем дна, нужно просто вернуть текущую строку и раскручивать стек вызовов обратно
	if brackets == n*2 {
		return []string{s}
	}
	var ss []string
	// чтобы поставить закрытую, нужно, чтобы была хоть одна открытая
	if opened > 0 {
		// количество закрытых не изменилось
		for _, sss := range a(s+")", n, brackets+1, opened-1, toClose) {
			ss = append(ss, sss)
		}
	}
	// чтобы поставить открытую, нужно, чтобы количество открытых не превышало максимум открытых-закрытых
	if toClose < n {
		for _, sss := range a(s+"(", n, brackets+1, opened+1, toClose+1) {
			ss = append(ss, sss)
		}
	}

	return ss
}
