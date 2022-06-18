package mergetwosortedlists

import (
	"testing"

	"gotest.tools/assert"
)

func TestMergeTwoShortLists(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}

	result := &ListNode{}

	assert.Equal(t, mergeTwoLists(list1, list2), result)
}
