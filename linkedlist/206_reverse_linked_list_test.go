package linkedlist

import (
	"fmt"
	"testing"
)

// 206. Reverse Linked List
// loop
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = pre
		pre = tmp
	}
	return pre
}

var newHead *ListNode

// recursive
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newTail := head
	reverseListHelper(head.Next).Next = head
	var pre *ListNode
	newTail.Next = pre
	return newHead
}

func reverseListHelper(head *ListNode) *ListNode {
	if head.Next == nil {
		newHead = head
		return head
	}
	reverseListHelper(head.Next).Next = head
	return head
}

func reverseList2020110701(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	tmp := head.Next
	for tmp != nil {
		head.Next = pre
		head, pre = tmp, head
		tmp = tmp.Next
	}
	head.Next = pre
	return head
}

func reverseList2020110702(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next, pre = pre, tmp
	}
	return pre
}

func reverseList2020110703(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := reverseList2020110703(head.Next)
	head.Next.Next = head
	head.Next = nil

	return pre
}

func reverseList20210113(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = pre
		pre = tmp
	}
	return pre
}

func TestReverseList(t *testing.T) {
	l := arrayToLinkedList([]int{1, 2, 3, 4})
	h := reverseList20210113(l)
	fmt.Println(h)
}
