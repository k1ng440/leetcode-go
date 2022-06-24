package lowestcommonancestorofabinarysearchtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	cur := root

	for {
		if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right
			continue
		}

		if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
			continue
		}

		break
	}

	return cur
}
