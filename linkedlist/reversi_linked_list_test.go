package linkedlist

import (
	"fmt"
	"testing"
)

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

func TestReverseList(t *testing.T) {
	l := arrayToLinkedList([]int{1})
	h := reverseList1(l)
	fmt.Println(h)
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
