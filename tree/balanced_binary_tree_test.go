package tree

import "math"

// 110. Balanced Binary Tree

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if math.Abs(float64(height(root.Left))-float64(height(root.Right))) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func isBalanced1(root *TreeNode) bool {
	res := true
	dfsWithLevelRecord(root, 1, &res)
	return res
}

func dfsWithLevelRecord(root *TreeNode, level int, res *bool) int {
	if root == nil {
		return level
	}
	if *res == false {
		return level
	}
	var lH, rH int
	if root.Left != nil {
		lH = dfsWithLevelRecord(root.Left, level+1, res)
		if *res == false {
			return level
		}
	} else {
		lH = level
	}
	if root.Right != nil {
		rH = dfsWithLevelRecord(root.Right, level+1, res)

	} else {
		rH = level
	}
	if math.Abs(float64(lH-rH)) > 1 {
		*res = false
	}
	return max(lH, rH)
}
