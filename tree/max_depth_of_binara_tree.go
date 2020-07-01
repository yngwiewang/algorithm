package tree

// leetcode 104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth1(root *TreeNode) int {
	var depth int
	if root == nil {
		return depth
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		tmp := []*TreeNode{root}
		for _, node := range stack {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		stack = tmp
		depth++
	}
	return depth
}
