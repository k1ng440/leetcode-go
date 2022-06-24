package implementqueueusingstacks

type MyQueue struct {
	stack []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	res := this.stack[0]
	this.stack = this.stack[1:]
	return res
}

func (this *MyQueue) Peek() int {
	return this.stack[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}
