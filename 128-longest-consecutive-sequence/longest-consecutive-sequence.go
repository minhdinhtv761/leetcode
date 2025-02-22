func longestConsecutive(nums []int) int {
	existMap := make(map[int]struct{})
	for _, num := range nums {
		existMap[num] = struct{}{}
	}
	res := 0
	for num := range existMap {
		if _, ok := existMap[num-1]; !ok {
			length := 1
			for _, ok1 := existMap[num+length]; ok1; _, ok1 = existMap[num+length] {
				length++
			}
			if length > res {
				res = length
			}
		}
	}

	return res
}