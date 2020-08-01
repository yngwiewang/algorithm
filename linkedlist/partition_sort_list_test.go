package linkedlist

import (
	"fmt"
	"testing"
)

func partitionSort(head *ListNode, x int) *ListNode {
	pSmall, pBig, pEqual := &ListNode{}, &ListNode{}, &ListNode{}
	p1, p2, p3 := pSmall, pBig, pEqual
	for head != nil {
		if head.Val < x {
			pSmall.Next = head
			pSmall = pSmall.Next
		} else if head.Val > x {
			pBig.Next = head
			pBig = pBig.Next
		} else {
			pEqual.Next = head
			pEqual = pEqual.Next
		}
		head = head.Next
	}
	pSmall.Next, pEqual.Next, pBig.Next = p3.Next, p2.Next, nil
	return p1
}

func TestPartitionSort(t *testing.T) {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  2,
							Next: nil,
						},
					},
				},
			},
		},
	}
	partitionSort(l, 3)
	head := l
	for head != nil {
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
}
