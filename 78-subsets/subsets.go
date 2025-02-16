func subsets(nums []int) [][]int {
	res := [][]int{[]int{}}

	var backtrack func(curIdx int, curArr []int)
	backtrack = func(curIdx int, curArr []int) {
        if len(curArr) == len(nums) || curIdx >= len(nums) {
			return
		}

		for idx := curIdx; idx < len(nums); idx++ {
            cl := make([]int, len(curArr))
            copy(cl, curArr)
			res = append(res, append(cl, nums[idx]))
            backtrack(idx+1, append(curArr, nums[idx]))
		}
	}

	backtrack(0, []int{})

	return res
}