func generateParenthesis(n int) []string {
	res := make([]string, 0)

	var backtrack func(open, close int, str string)
	backtrack = func(open, close int, str string) {
		if len(str) == 2*n {
			res = append(res, str)
			return
		}

		if open < n {
			backtrack(open+1, close, str + "(")
		}

		if close < open {
			backtrack(open, close+1, str + ")")
		}
	}

	backtrack(0, 0, "")
	return res
}