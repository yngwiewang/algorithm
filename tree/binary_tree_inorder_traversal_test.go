package tree

import "testing"

// 94. Binary Tree Inorder Traversal

// loop
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for ; root != nil; root = root.Left {
			stack = append(stack, root)
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			root = node.Right
		}
	}
	return res
}

// recursive
func inorderTraversal1(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)

	return res
}

func TestInorderTraversal1(t *testing.T) {
	tree := arrayToBinaryTree([]int{1, 2, 3, 4, 5, 6, 7, 9})
	t.Log(inorderTraversal1(tree))
}
