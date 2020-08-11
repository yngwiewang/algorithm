package stack

import "testing"

//225. Implement Stack using Queues

type MyStack struct {
	q1 []int
	q2 []int
}

/** Initialize your data structure here. */
func Constructor1() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if this.Empty() {
		this.q1 = append(this.q1, x)
		return
	}
	if this.q1 == nil {
		this.q1, this.q2 = this.q2, this.q1
	}
	this.q2 = append(this.q2, x)
	for _, v := range this.q1 {
		this.q2 = append(this.q2, v)
	}
	this.q1 = nil
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.q1 == nil {
		this.q1, this.q2 = this.q2, this.q1
	}
	this.q2 = nil
	pop := this.q1[0]

	if len(this.q1) > 1 {
		this.q1 = this.q1[1:]
	} else {
		this.q1 = nil
	}
	return pop
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.q1 == nil {
		this.q1, this.q2 = this.q2, this.q1
	}
	return this.q1[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.q1 == nil && this.q2 == nil {
		return true
	}
	return false
}

func Test_StackUsingQueues(t *testing.T) {
	s := Constructor1()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	t.Log(s.Top())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Empty())
	t.Log(s.Pop())
	t.Log(s.Empty())
}
