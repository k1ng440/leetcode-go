package maximumsubarray

import "math"

func maxSubArray(nums []int) int {
	sum := int(-1e5)
	max := int(-1e5)
	for _, num := range nums {
		sum = int(math.Max(float64(num), float64(sum+num)))
		max = int(math.Max(float64(max), float64(sum)))
	}

	return max
}
