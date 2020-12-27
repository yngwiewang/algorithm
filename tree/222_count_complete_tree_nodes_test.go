package tree

// 222. Count Complete Tree Nodes
// didn't use binary search
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func countNodes11(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		res += size
		for _, v := range q {
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}
		q = q[size:]
	}
	return res
}
