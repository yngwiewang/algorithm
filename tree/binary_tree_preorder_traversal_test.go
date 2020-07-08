package tree

import (
	"bytes"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode 144

func preorderTraversal(root *TreeNode) []int {
	var pre []int
	if root == nil {
		return pre
	}
	pre = append(pre, root.Val)
	pre = append(pre, preorderTraversal(root.Left)...)
	pre = append(pre, preorderTraversal(root.Right)...)
	return pre
}

func preorderTraversalLoop(root *TreeNode) []int {
	var pre []int
	if root == nil {
		return pre
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		pre = append(pre, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return pre
}
