package groupanagrams

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, str := range strs {
		x := [26]int{}
		for _, s := range str {
			x[s-'a']++
		}

		groups[x] = append(groups[x], str)
	}

	res := make([][]string, 0)
	for _, group := range groups {
		res = append(res, group)
	}

	return res
}
