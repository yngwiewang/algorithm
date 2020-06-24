package linkedlist

import (
	"fmt"
	"testing"
)

// leetcode 92

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	p0 := &ListNode{0, nil}
	p0.Next = head
	p := p0
	i := 1
	for ; i < m; i++ {
		p0 = p0.Next
	}
	c := p0.Next
	for ; i < n; i++ {
		a, b := c.Next, c.Next.Next
		a.Next = p0.Next
		c.Next = b
		p0.Next = a
	}

	return p.Next
}

func TestReverseBetween(t *testing.T) {
	a := &ListNode{}
	p := a
	for i := 1; i < 6; i++ {
		a.Next = &ListNode{i, nil}
		a = a.Next
	}
	x := p.Next

	y := reverseBetween(x, 3, 5)
	for y != nil {
		fmt.Printf("%d ", y.Val)
		y = y.Next
	}
	fmt.Println()
}
