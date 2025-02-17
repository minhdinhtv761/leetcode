func exist(board [][]byte, word string) bool {
	mark := make([][]int, len(board))
	for i := range mark {
		mark[i] = make([]int, len(board[0]))
	}

	var dfs func(r, c, i int) bool
	dfs = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}

		if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) || mark[r][c] == 1 || board[r][c] != word[i] {
			return false
		}

		mark[r][c] = 1
		res := dfs(r+1, c, i+1) || dfs(r-1, c, i+1) || dfs(r, c+1, i+1) || dfs(r, c-1, i+1)
		mark[r][c] = 0
        return res
	}

    for r := range board {
        for c := range board[r] {
            res := dfs(r, c, 0)
            if res {
                return true
            }
        }
    }

	return false
}