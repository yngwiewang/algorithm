package linkedlist

import (
	"fmt"
	"testing"
)

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			goto REACH
		}
	}
	return nil
REACH:
	for head != slow {
		head, slow = head.Next, slow.Next
	}
	return head
}

func TestDetectCycle(t *testing.T) {
	a, b, c := &ListNode{1, nil}, &ListNode{2, nil}, &ListNode{3, nil}
	a.Next, b.Next, c.Next = b, c, a
	res := detectCycle(a)
	fmt.Println(res)
}
