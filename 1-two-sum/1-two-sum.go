package twosum

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for rIdx, rNum := range nums {
		lIdx, ok := seen[rNum-target]
		if !ok {
			seen[rNum] = rIdx
			continue
		}

		return []int{rIdx, lIdx}
	}

	return nil
}
