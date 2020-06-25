package linkedlist

import (
	"fmt"
	"testing"
)

// leetcode 82

func deleteDuplicates2(head *ListNode) *ListNode {
	p0 := &ListNode{
		Val:  0,
		Next: head,
	}
	pPre := p0
	for head != nil {
		var dup bool
		for head.Next != nil && head.Val == head.Next.Val {
			head.Next = head.Next.Next
			dup = true
		}
		if dup == true {
			head = head.Next
			pPre.Next = head
		} else {
			pPre.Next = head
			pPre = pPre.Next
			head = pPre.Next
		}
	}
	return p0.Next
}

func TestDeleteDuplicates2(t *testing.T) {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	res := deleteDuplicates2(l)
	for res != nil {
		fmt.Printf("%v ", res.Val)
		res = res.Next
	}
}
