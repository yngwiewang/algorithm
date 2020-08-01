package linkedlist

import (
	"testing"
)

// 141. Linked List Cycle
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func TestHasCycle(t *testing.T) {
	n := &ListNode{1, nil}
	n.Next = n
	t.Log(hasCycle(n))
}
