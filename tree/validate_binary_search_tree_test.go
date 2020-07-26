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

// recursive divide and conquer
func isValidBST3(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, _, ok := isValidBSTHelper1(root)
	return ok
}

func isValidBSTHelper1(root *TreeNode) (l, r int, res bool) {
	res = true
	if root.Left == nil && root.Right == nil {
		return root.Val, root.Val, true
	}
	lMin, lMax, left := math.MinInt64, math.MinInt64, true
	rMin, rMax, right := math.MaxInt64, math.MaxInt64, true
	if root.Left != nil {
		lMin, lMax, left = isValidBSTHelper1(root.Left)
	}
	m := root.Val
	if root.Right != nil {
		rMin, rMax, right = isValidBSTHelper1(root.Right)
	}

	if lMax >= m || m >= rMin {
		res = false
	} else {
		if root.Left == nil {
			l = m
		} else {
			l = lMin
		}
		if root.Right == nil {
			r = m
		} else {
			r = rMax
		}
	}
	res = res && left && right
	return
}
