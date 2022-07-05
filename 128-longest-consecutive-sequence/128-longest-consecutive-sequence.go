package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	max := 0
	mp := make(map[int]int)

	for _, n := range nums {
		mp[n] = 1
	}

	for i, n := range nums {
		if _, ok := mp[i-1]; ok {
			mp[n] = 0
		}
	}

outtaLoop:
	for n, x := range mp {
		if x != 1 {
			continue
		}

		curr := 1

		for {
			if _, ok := mp[n+curr]; ok {
				curr++
				continue
			}

			if max < curr {
				max = curr
			}

			continue outtaLoop
		}
	}

	return max
}
