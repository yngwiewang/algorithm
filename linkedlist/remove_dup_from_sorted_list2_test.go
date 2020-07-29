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
			pPre = head
			head = head.Next
		}
	}
	return p0.Next
}

func TestDeleteDuplicates2(t *testing.T) {
	l := arrayToLinkedList([]int{1, 1, 1, 1, 2, 2, 2, 5, 6, 7, 7, 7})

	res := deleteDuplicates2(l)
	for res != nil {
		fmt.Printf("%v ", res.Val)
		res = res.Next
	}
}
