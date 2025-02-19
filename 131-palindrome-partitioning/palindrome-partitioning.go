func partition(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)

	var dfs func(i int)
	dfs = func(i int) {
		if i >= len(s) {
			clone := make([]string,	len(path))
			copy(clone, path)
			res = append(res, clone)
		}

		for j := i; j < len(s); j++ {
			if isPalindrome(s, i, j) {
				path = append(path, string(s[i:j+1]))
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)

	return res
}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}