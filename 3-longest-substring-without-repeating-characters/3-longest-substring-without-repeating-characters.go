package leetcode

func lengthOfLongestSubstring(s string) int {
	chars := make(map[rune]int)
	runes := []rune(s)
	res := 0
	left := 0

	for right, x := range runes {
		if _, ok := chars[x]; ok {
			left = max(chars[x], left)
		}

		res = max(res, (right-left)+1)
		chars[x] = right + 1
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
