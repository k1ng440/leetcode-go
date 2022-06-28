package convertsortedarraytobinarysearchtree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return solve(nums, 0, len(nums)-1)
}

func solve(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := int(math.Floor(float64(left + (right-left)/2)))

	return &TreeNode{
		Val:   nums[mid],
		Left:  solve(nums, left, mid-1),
		Right: solve(nums, mid+1, right),
	}
}
