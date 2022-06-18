package binarysearch

func search(nums []int, target int) int {
	return binarysearch(nums, target, 0, len(nums)-1)
}

func binarysearch(nums []int, target, left, right int) int {
	if left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			return binarysearch(nums, target, left, mid-1)
		}

		return binarysearch(nums, target, mid+1, right)
	}

	return -1
}
