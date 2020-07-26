package tree

import (
	"math"
	"testing"
)

func isValidBST(root *TreeNode) bool {
	inorder := inorderTraversal(root)
	validate := true
	for i := 1; i < len(inorder); i++ {
		if inorder[i] <= inorder[i-1] {
			validate = false
			break
		}
	}
	return validate
}

func TestIsValidBST(t *testing.T) {
	tree := arrayToBinaryTree([]int{2, 1, 3})
	t.Log(isValidBST1(tree))
}

// less mem
func isValidBST1(root *TreeNode) bool {
	_, validate := inorderTraversal_(root)
	return validate
}

func inorderTraversal_(root *TreeNode) (pre int, res bool) {
	res = true
	stack := []*TreeNode{}
	afterOnePass := false
OUTER:
	for root != nil || len(stack) > 0 {
		for ; root != nil; root = root.Left {
			stack = append(stack, root)
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if afterOnePass && node.Val <= pre {
			res = false
			break OUTER
		}
		pre = node.Val
		afterOnePass = true
		if node.Right != nil {
			root = node.Right
		}
	}
	return
}

func isValidBST2(root *TreeNode) bool {
	return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTHelper(root *TreeNode, low, higt int) bool {
	if root == nil {
		return true
	}
	if root.Val <= low || root.Val >= higt {
		return false
	}
	return isValidBSTHelper(root.Left, low, root.Val) &&
		isValidBSTHelper(root.Right, root.Val, higt)
}

func TestisValidBSTHelper(t *testing.T) {
	tree := arrayToBinaryTree([]int{1, 1})
	t.Log(isValidBSTHelper(tree, math.MinInt32, math.MaxInt32))
}
