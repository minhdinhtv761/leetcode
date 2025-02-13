func combinationSum(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool { return candidates[i] < candidates[j] })

	res := make([][]int, 0)

	var backtrack func(idx, curSum int, curCombination []int)
	backtrack = func(idx, curSum int, curCombination []int) {
		if curSum == target {
            clone := make([]int, len(curCombination))
            copy(clone, curCombination)
			res = append(res, clone)
			return
		}

		if curSum > target || idx >= len(candidates) {
			return
		}

		backtrack(idx, curSum+candidates[idx], append(curCombination, candidates[idx]))
		backtrack(idx+1, curSum, curCombination)
	}

	backtrack(0, 0, []int{})
	return res
}