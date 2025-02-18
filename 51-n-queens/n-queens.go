func solveNQueens(n int) [][]string {
	mark := make(map[int]int)
	res := make([][]string, 0)

	var backtrack func(c int)
	backtrack = func(c int) {
		if c == n {
			pos := make([]string, 0)
			for r := range n {
				pos = append(pos, replaceAtIndex(strings.Repeat(".", n), 'Q', mark[r]))
			}
			res = append(res, pos)
			return
		}
		for r := range n {
			if _, ok := mark[r]; ok {
				continue
			}
			tmp := false
			for r1 := range n {
				if c1, ok := mark[r1]; ok && (r-r1 == c-c1 || r-r1 == (c-c1) * -1) {
					tmp = true
					break
				}
			}
			if tmp {
				continue
			}
			mark[r] = c
			backtrack(c + 1)
			delete(mark, r)
		}
	}

	backtrack(0)
	return res
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}