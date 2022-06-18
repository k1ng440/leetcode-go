package leetcode

import "fmt"

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	bucket := make(map[int]int)
	width := t + 1

	for i, num := range nums {
		j := (num / width) - 1

		if num >= 0 {
			j = num / width
		}

		if _, exist := bucket[j]; exist {
			return true
		} else if bucketNum, exist := bucket[j+1]; exist && (abs(num-bucketNum) < width) {
			return true
		} else if bucketNum, exist := bucket[j-1]; exist && (abs(num-bucketNum) < width) {
			return true
		}

		// put current num into corresponding bucket
		bucket[j] = num

		if i >= k {
			fmt.Println(nums[i-k] / width)
			delete(bucket, nums[i-k]/width)
		}
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
