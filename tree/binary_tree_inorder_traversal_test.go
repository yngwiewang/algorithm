package tree

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
