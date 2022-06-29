package implementstackusingqueues

type MyStack struct {
	head *Node
}

type Node struct {
	Val  int
	Next *Node
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	curr := this.head
	this.head = &Node{Val: x, Next: curr}
}

func (this *MyStack) Pop() int {
	val := this.head.Val
	this.head = this.head.Next
	return val
}

func (this *MyStack) Top() int {
	return this.head.Val
}

func (this *MyStack) Empty() bool {
	return this.head == nil
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
