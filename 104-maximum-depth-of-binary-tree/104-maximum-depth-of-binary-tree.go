package maximumdepthofbinarytree

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return int(dfs(root))
}

func dfs(root *TreeNode) float64 {
	if root == nil {
		return 0
	}

	return 1 + math.Max(dfs(root.Left), dfs(root.Right))
}
