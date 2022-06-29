package searchinsertposition

func searchInsert(nums []int, target int) int {
	return search(nums, target, 0, len(nums)-1)
}

func search(nums []int, target, left, right int) int {
	if left > right {
		return left
	}

	mid := left + (right-left)/2

	if mid >= len(nums) {
		return len(nums)
	}

	if target == nums[mid] {
		return mid
	}

	if target < nums[mid] {
		return search(nums, target, left, mid-1)
	}

	return search(nums, target, mid+1, right)
}
