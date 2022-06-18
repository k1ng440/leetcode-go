package leetcode

func maxArea(height []int) int {
	ma := 0
	l := 0
	r := len(height) - 1

	for l < r {
		ma = max(ma, min(height[l], height[r])*(r-l))

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return ma
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
