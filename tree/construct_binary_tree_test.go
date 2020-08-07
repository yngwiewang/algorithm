package tree

import (
	"testing"
)

// 105. Construct Binary Tree from Preorder and Inorder Traversal
func buildTreeByPreIn(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	p := 0
	for ; p < len(inorder); p++ {
		if inorder[p] == preorder[0] {
			break
		}
	}
	root.Left = buildTreeByPreIn(preorder[1:1+p], inorder[0:p])
	root.Right = buildTreeByPreIn(preorder[1+p:], inorder[p+1:])
	return root
}

func Test_buildTree(t *testing.T) {
	pre := []int{1, 2, 4, 5, 3, 6, 7}
	in := []int{4, 2, 5, 1, 6, 3, 7}
	tree := buildTreeByPreIn(pre, in)
	newPre := preorderTraversal(tree)
	newIn := inorderTraversal(tree)
	t.Log(newPre)
	t.Log(newIn)
}
