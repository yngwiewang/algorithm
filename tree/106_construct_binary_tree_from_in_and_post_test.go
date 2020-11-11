package tree

import "testing"

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

// review 20201111
func buildTreeByInPost20201111(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			break
		}
	}
	return &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  buildTreeByInPost20201111(inorder[:i], postorder[:i]),
		Right: buildTreeByInPost20201111(inorder[i+1:], postorder[i:len(postorder)-1]),
	}
}
func Test_buildTreeByInPost(t *testing.T) {
	in := []int{4, 2, 5, 1, 6, 3, 7}
	post := []int{4, 5, 2, 6, 7, 3, 1}
	tree := buildTreeByInPost20201111(in, post)
	newPre := preorderTraversal(tree)
	newIn := inorderTraversal(tree)
	newPost := postorderTraversal(tree)
	t.Log(newPre)
	t.Log(newIn)
	t.Log(newPost)
}
