package linkedlist

import (
	"fmt"
	"testing"
)

// not stable
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

// stable
func inser(head *ListNode) *ListNode {
	pre := &ListNode{0, head}
	for head != nil && head.Next != nil {
		for ; head.Next != nil && head.Val <= head.Next.Val; head = head.Next {
		}
		if head.Next == nil {
			break
		}
		cur := head.Next
		next := cur.Next
		tmp := pre
		for ; cur.Val > tmp.Next.Val; tmp = tmp.Next {
		}
		head.Next = next
		cur.Next = tmp.Next
		tmp.Next = cur
	}
	return pre.Next
}

func TestInsertionSortList(t *testing.T) {
	l := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	//s := insertionSortList(l)
	s := inser(l)
	for s != nil {
		fmt.Printf("%d ", s.Val)
		s = s.Next
	}
}
