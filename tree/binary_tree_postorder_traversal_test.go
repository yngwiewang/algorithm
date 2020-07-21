package tree

// leetcode 145

// recursive

func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	if root.Left != nil {
		res = append(res, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, postorderTraversal(root.Right)...)
	}

	res = append(res, root.Val)
	return res
}

// 前序遍历后翻转

func postorderTraversal2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

// 循环

func postorderTraversal3(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := []*TreeNode{}
	var lastVisit *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return res
}
