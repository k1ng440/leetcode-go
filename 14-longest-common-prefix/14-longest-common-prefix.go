package longestcommonprefix

import "sort"

func longestCommonPrefix(strs []string) string {
	// sort by length of string to get the sortest
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})

	tmp := make([]byte, 0)
	for i := 0; i < len(strs[0]); i++ {
		lit := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if lit != strs[j][i] {
				return ""
			}
		}

		tmp = append(tmp, lit)
	}

	return string(tmp)
}
