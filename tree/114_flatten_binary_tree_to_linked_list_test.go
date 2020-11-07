package tree

import "testing"

// 114. Flatten Binary Tree to Linked List
// 1 preorderTraversal, 2 iterate
func flattenA(root *TreeNode) {
	if root == nil {
		return
	}
	preorder := []*TreeNode{}
	preorderNodes(root, &preorder)

	for i := len(preorder) - 2; i >= 0; i-- {
		preorder[i].Right = preorder[i+1]
		preorder[i].Left = nil
	}
	// tmp := &TreeNode{}
	// for _, node := range preorder{
	// 	tmp.Right = node
	// 	tmp.Left = nil
	// 	tmp = tmp.Right
	// }
}

func preorderNodes(root *TreeNode, res *[]*TreeNode) {
	if root == nil {
		return
	}
	*res = append(*res, root)
	preorderNodes(root.Left, res)
	preorderNodes(root.Right, res)
	return
}

// recursive,beat 100%
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	r := root.Right
	root.Right = root.Left
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = r
}

func Test_flatten(t *testing.T) {
	a := arrayToBinaryTree([]int{1, 2, 3, 4, 5, 6})
	flattenA(a)
	t.Log(preorderTraversal(a))
}
