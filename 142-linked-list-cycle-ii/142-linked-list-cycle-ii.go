package linkedlistcycleii

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	meet := false

	for fast != nil && fast.Next != nil && meet == false {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			meet = true
		}
	}

	if meet == false || fast == nil {
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
