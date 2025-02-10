var m = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	if len(digits) == 1 {
		return m[digits[0]]
	}
	sub := letterCombinations(digits[1:])
	for _, s := range m[digits[0]] {
		for _, ss := range sub {
			res = append(res, fmt.Sprintf("%v%v", s, ss))
		}
	}
	return res
}