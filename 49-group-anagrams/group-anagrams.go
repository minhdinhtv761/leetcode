func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[string][]string)

	for _, str := range strs {
		count := make(map[int32]int)
		for _, c := range str {
			count[c]++
		}
		key := buildKey(count)
		if _, ok := hashMap[key]; !ok {
			hashMap[key] = make([]string, 0)
		}
		hashMap[key] = append(hashMap[key], str)
	}

    res := make([][]string, 0)
    for _, v := range hashMap {
        res = append(res, v)
    }

    return res
}

func buildKey(m map[int32]int) string {
	key := ""
	for i := 'a'; i <= 'z'; i++ {
		if v, ok := m[i]; ok {
			key += fmt.Sprintf("%d%s", v, string(i))
		}
	}
	return key
}