package linkedlist

import (
	"fmt"
	"testing"
)

// leetcode 109

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{head.Val, nil, nil}
	}
	slow, fast := head, head
	preSlow := &ListNode{0, head}
	for fast != nil && fast.Next != nil {
		preSlow, slow, fast = preSlow.Next, slow.Next, fast.Next.Next
	}
	root := &TreeNode{slow.Val, nil, nil}
	preSlow.Next = nil
	root.Left = sortedListToBST(head)
	if slow.Next != nil {
		root.Right = sortedListToBST(slow.Next)
	}
	return root
}

func TestSortedListToBST(t *testing.T) {
	a := arrayToLinkedList([]int{1, 2, 3, 4, 5})
	b := sortedListToBST(a)
	fmt.Println(b)
}
