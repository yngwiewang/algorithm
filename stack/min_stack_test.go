package minstack

import (
	"testing"
)

// 155. Min Stack

type MinStack struct {
	stack [][2]int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([][2]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, [2]int{x, x})
	} else {
		mini := min(x, this.stack[len(this.stack)-1][1])
		this.stack = append(this.stack, [2]int{x, mini})
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1][0]
}

func (this *MinStack) GetMin() int {	
	return this.stack[len(this.stack)-1][1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Test_minStack(t *testing.T) {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)	
	t.Log(s)

	
	t.Log(s.GetMin())

	s.Pop()
	
	t.Log(s.Top())
	
	t.Log(s.GetMin())

}
