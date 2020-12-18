package tree

import "testing"

// 226. Invert Binary Tree
// recursive, postorder
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func invertTreePreorder(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)
	return root
}

// iterate
func invertTreeA(root *TreeNode) *TreeNode {
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q[0]
		q = q[1:]
		if tmp == nil || (tmp.Left == nil && tmp.Right == nil) {
			continue
		}
		tmp.Left, tmp.Right = tmp.Right, tmp.Left
		q = append(q, tmp.Right)
		q = append(q, tmp.Left)

	}
	return root
}

func Test_invertTree(t *testing.T) {
	a := arrayToBinaryTree([]int{1, 2})
	invertTreePreorder(a)

}
