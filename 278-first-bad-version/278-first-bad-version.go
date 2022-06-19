package firstbadversion

// this function is provided by leetcode
func isBadVersion(int) bool { return false }

func firstBadVersion(n int) int {
	left := 0
	right := n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}
