package besttimetobuyandsellstock

import "math"

func maxProfit(prices []int) int {
	l := 0
	maxPrice := 0

	for r := 1; r < len(prices); r++ {
		if prices[r] < prices[l] {
			l = r
		}

		maxPrice = int(math.Max(float64(maxPrice), float64(prices[r]-prices[l])))
	}

	return maxPrice
}
