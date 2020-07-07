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

var head *ListNode

func convert(l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	left := convert(l, mid-1)
	node := &TreeNode{head.Val, nil, nil}
	node.Left = left
	head = head.Next
	node.Right = convert(mid+1, r)
	return node
}

func sortedListToBST1(head *ListNode) *TreeNode {
	c := 0
	for n := head; n != nil; n = n.Next {
		c++
	}
	return convert(0, c-1)
}

func TestSortedListToBST(t *testing.T) {
	a := arrayToLinkedList([]int{1, 2, 3, 4, 5})
	b := sortedListToBST(a)
	fmt.Println(b)
}

// inorder traversal, slower

//var head *ListNode
//
//func sortedListToBST(h *ListNode) *TreeNode {
//	var s int
//	head = h
//	for h != nil {
//		s++
//		h = h.Next
//	}
//	return convertListToBST(0, s-1)
//}
//
//func convertListToBST(l, r int) *TreeNode {
//	if l > r {
//		return nil
//	}
//	mid := (l + r) / 2
//	left := convertListToBST(l, mid-1)
//	node := &TreeNode{
//		Val: head.Val,
//	}
//	node.Left = left
//	head = head.Next
//	node.Right = convertListToBST(mid+1, r)
//	return node
//}
