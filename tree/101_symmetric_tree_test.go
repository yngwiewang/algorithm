package tree

// 101. Symmetric Tree

// recursive with two nodes
func isSymmetric(root *TreeNode) bool {

	return isSymmetricHelper(root, root)
}

func isSymmetricHelper(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.Val == r.Val && isSymmetricHelper(l.Left, r.Right) && isSymmetricHelper(l.Right, r.Left)
}
