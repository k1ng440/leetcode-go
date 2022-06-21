package minstack

import "math"

type node struct {
	val int
	min int
}

type MinStack struct {
	nodes []*node
}

func Constructor() MinStack {
	return MinStack{
		nodes: make([]*node, 0),
	}
}

func (this *MinStack) Push(val int) {
	min := val
	if l := len(this.nodes) - 1; l != -1 {
		lastNode := this.nodes[l]
		min = int(math.Min(float64(lastNode.min), float64(val)))
	}
	this.nodes = append(this.nodes, &node{
		val: val,
		min: min,
	})
}

func (this *MinStack) Pop() {
	this.nodes = this.nodes[:len(this.nodes)-1]
}

func (this *MinStack) Top() int {
	return this.nodes[len(this.nodes)-1].val
}

func (this *MinStack) GetMin() int {
	return this.nodes[len(this.nodes)-1].min
}
