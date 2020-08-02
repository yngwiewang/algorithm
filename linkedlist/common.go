package linkedlist

import (
	"strconv"
)

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func arrayToLinkedList(a []int) *ListNode {
	l := new(ListNode)
	pre := l
	for _, i := range a {
		l.Next = &ListNode{i, nil}
		l = l.Next
	}
	return pre.Next
}

func (l *ListNode) String() string {
	var res string
	for l != nil {
		res += " " + strconv.Itoa(l.Val)
		l = l.Next
	}
	return res
}
