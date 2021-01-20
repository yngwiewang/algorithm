package tree

// 235. Lowest Common Ancestor of a Binary Search Tree
func lowestCommonAncestorInBST(root, p, q *TreeNode) *TreeNode {
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val < root.Val && q.Val > root.Val || p.Val > root.Val && q.Val < root.Val {
		return root
	}
	if p.Val == root.Val {
		return p
	}
	return q
}
