package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	head := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l2
			l2 = l2.Next
			continue
		} else {
			head.Next = l1
			l1 = l1.Next
		}

		head = head.Next
	}

	if l1 != nil {
		head.Next = l1
	}

	if l2 != nil {
		head.Next = l2
	}

	return dummy.Next
}
