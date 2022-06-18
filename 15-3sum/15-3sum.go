package sum3

import (
	"sort"
)

// threeSum([]int{-1, 0, 1, 2, -1, -4}) = [	[-1, 0, 1], [-1, -1, 2] ]
// threeSum([]int{0, 0, 0, 0}) = [	[0, 0, 0] ]

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	sort.IntSlice(nums).Sort()
	for i, a := range nums {
		if i > 0 && a == nums[i-1] {
			continue
		} 
        
		l, r := i+1, len(nums)-1
		for l < r {
			ts := a + nums[l] + nums[r]
			if ts > 0 {
                r--
			} else if ts < 0 {
				l++
			} else {
				ret = append(ret, []int{a, nums[l], nums[r]})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}
	return ret
}
