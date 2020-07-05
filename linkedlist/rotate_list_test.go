package linkedlist

import (
	"fmt"
	"strings"
	"testing"
)

func assembleLinkedList(l []*ListNode) *ListNode {
	head := new(ListNode)
	pre := head
	for _, n := range l {
		head.Next = n
		head = n
	}
	return pre.Next
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var list []*ListNode
	for ; head != nil; head = head.Next {
		list = append(list, &ListNode{head.Val, nil})
	}
	k = k % len(list)
	for i := 0; i < k; i++ {
		list = append(list[len(list)-1:], list[0:len(list)-1]...)
	}
	return assembleLinkedList(list)
}

func rotateRight1(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var list []*ListNode
	for ; head != nil; head = head.Next {
		list = append(list, &ListNode{head.Val, nil})
	}
	k = k % len(list)
	for i := 0; i < k; i++ {
		tmp := list[len(list)-1]
		copy(list[1:], list[:len(list)-1])
		list[0] = tmp
	}
	return assembleLinkedList(list)
}

func rotateRight2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	length := 1
	p := head
	for ; p.Next != nil; p = p.Next {
		length++
	}
	k = k % length
	if k == 0 {
		return head
	}
	oldHead := head
	for i := 1; i < length-k; i++ {
		head = head.Next
	}
	newHead := head.Next
	head.Next = nil
	p.Next = oldHead
	return newHead
}

func TestRotateList(t *testing.T) {
	a := arrayToLinkedList([]int{1, 2, 3, 4, 5})
	l := rotateRight2(a, 5)
	fmt.Println(l)
}

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}
	var res []string
	for ; l != nil; l = l.Next {
		res = append(res, fmt.Sprintf("%d", l.Val))
	}
	return strings.Join(res, " ")
}
