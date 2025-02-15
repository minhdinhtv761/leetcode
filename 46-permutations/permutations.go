func permute(nums []int) [][]int {
    res := make([][]int, 0)
    mark := make([]int, len(nums))

    var backtrack func(cur []int)
    backtrack = func(cur []int) {
        if len(cur) == len(nums) {
            clone := make([]int, len(cur))
            copy(clone, cur)
            res = append(res, clone)
            return
        }

        for idx := range mark {
            if mark[idx] == 1 {
                continue
            }
            mark[idx] = 1
            backtrack(append(cur, nums[idx]))
            mark[idx] = 0
        }
    }

    backtrack([]int{})
    return res
}