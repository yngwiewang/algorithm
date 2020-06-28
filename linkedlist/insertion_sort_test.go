package linkedlist

import (
	"fmt"
	"testing"
)

func insertionSortList(head *ListNode) *ListNode {
	p0 := &ListNode{
		Val:  0,
		Next: head,
	}

	for head != nil && head.Next != nil {
		curr := head.Next
		next := curr.Next
		temp := p0
		for ; curr.Val > temp.Next.Val; temp = temp.Next {
		}
		if temp.Next != curr {
			head.Next = next
			curr.Next = temp.Next
			temp.Next = curr
		}
		for head.Next != nil && head.Val < head.Next.Val {
			head = head.Next
		}

	}
	return p0.Next
}

func TestInsertionSortList(t *testing.T) {
	l := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	s := insertionSortList(l)
	for s != nil{
		fmt.Printf("%d ", s.Val)
		s = s.Next
	}
}
