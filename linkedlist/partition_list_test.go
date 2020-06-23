package linkedlist

import (
	"fmt"
	"testing"
)

// leetcode 86

func partition(head *ListNode, x int) *ListNode {
	pSmall, pBig := new(ListNode), new(ListNode)
	p1, p2 := pSmall, pBig
	for head != nil {
		if head.Val < x {
			pSmall.Next = head
			pSmall = pSmall.Next
		} else {
			pBig.Next = head
			pBig = pBig.Next
		}
		head = head.Next
	}
	pSmall.Next, pBig.Next = p2.Next, nil
	return p1.Next
}

func TestPartition(t *testing.T) {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  2,
							Next: nil,
						},
					},
				},
			},
		},
	}
	partition(l, 3)
	head := l
	for head != nil {
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
}
