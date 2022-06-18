package coinchange

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				if dp[i] > dp[i-coin]+1 {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}
