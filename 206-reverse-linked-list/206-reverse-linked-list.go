package reverselinkedlist

type LinkedList struct {
    Val int
    Next *LinkedList
}

func reverseList(head *LinkedList) *LinkedList {
    var (
        prev *LinkedList 
        curr = head
    )

    for curr != nil {
        temp := head.Next
        curr.Next = prev
        prev = curr
        curr = temp
    }

    return prev
}

