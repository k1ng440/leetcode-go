package leetcode

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	runes := []rune(s)
	start := 0
	end := 0

	for i := range runes {
		maxLen := max(expendMiddle(runes, i, i), expendMiddle(runes, i, i+1))
		if maxLen > (end - start) {
			start = i - ((maxLen - 1) / 2)
			end = i + (maxLen / 2) + 1
		}
	}

	return string(runes[start:end])
}

func expendMiddle(runes []rune, left, right int) int {
	if len(runes) == 0 || left > right {
		return 0
	}

	for left >= 0 && right < len(runes) && runes[left] == runes[right] {
		left--
		right++
	}

	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
