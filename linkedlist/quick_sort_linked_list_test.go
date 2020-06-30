package linkedlist

import (
	"fmt"
	"testing"
)

// leetcode 148

func partitionSort(head, tail *ListNode) *ListNode {
	if head == tail {
		return head
	}

	pivot, head := head, head.Next
	small, big := &ListNode{0, nil}, &ListNode{0, nil}
	preSmall, preBig := &ListNode{0, small}, &ListNode{0, big}
	var moveSmall, moveBig bool
	for ; head != tail.Next; head = head.Next {
		if head.Val >= pivot.Val {
			big.Next = head
			big = big.Next
			moveBig = true
		} else {
			small.Next = head
			small = small.Next
			moveSmall = true
		}
	}
	small.Next, pivot.Next, big.Next = nil, nil, nil
	if moveSmall {
		preSmall.Next.Next = partitionSort(preSmall.Next.Next, small)
		for ; small.Next != nil; small = small.Next {
		}
	}
	if moveBig {
		preBig.Next.Next = partitionSort(preBig.Next.Next, big)
	}
	small.Next, pivot.Next = pivot, preBig.Next.Next
	return preSmall.Next.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre := &ListNode{0, head}
	for head.Next != nil {
		head = head.Next
	}
	tail := head
	res := partitionSort(pre.Next, tail)
	return res
}

func TestQuickSortLinkedList(t *testing.T) {
	s := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	s = &ListNode{
		Val:  2,
		Next: &ListNode{1, nil}}
	s = nil
	l := sortList(s)
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
}
