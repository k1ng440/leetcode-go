package reverselinkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var (
		curr = head
		prev *ListNode
	)

	for curr != nil {
		realNext := curr.Next
		curr.Next = prev
		prev = curr
		curr = realNext
	}

	return curr
}
