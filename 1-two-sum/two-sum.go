func twoSum(nums []int, target int) []int {
	check := make(map[int]int)
    for i := range nums {
        if v, ok := check[target-nums[i]]; ok {
            return []int{v, i}
        }
        check[nums[i]] = i 
	}
	return nil
}