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

// 106. Construct Binary Tree from Inorder and Postorder Traversal
func buildTreeByInPost(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}
	p := 0
	root := &TreeNode{Val: postorder[(len(postorder) - 1)]}
	for ; p < len(inorder); p++ {
		if inorder[p] == postorder[(len(postorder)-1)] {
			break
		}
	}
	root.Left = buildTreeByInPost(inorder[:p], postorder[:p])
	root.Right = buildTreeByInPost(inorder[p+1:], postorder[p:(len(postorder)-1)])
	return root
}

func Test_buildTreeByInPost(t *testing.T) {
	in := []int{4, 2, 5, 1, 6, 3, 7}
	post := []int{4, 5, 2, 6, 7, 3, 1}
	tree := buildTreeByInPost(in, post)
	newPre := preorderTraversal(tree)
	newIn := inorderTraversal(tree)
	t.Log(newPre)
	t.Log(newIn)
}
