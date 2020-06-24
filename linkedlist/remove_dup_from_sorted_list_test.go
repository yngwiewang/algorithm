package linkedlist

import (
	"fmt"
	"testing"
)

// leetcode 83

func deleteDuplicates(head *ListNode) *ListNode {
	p0 := head
	for head != nil {
		for head.Next != nil && head.Next.Val == head.Val {
			head.Next = head.Next.Next
		}
		head = head.Next
	}
	return p0
}

func TestDeleteDuplicates(t *testing.T) {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}

	res := deleteDuplicates(l)
	for res != nil {
		fmt.Printf("%v ", res.Val)
		res = res.Next
	}
}
