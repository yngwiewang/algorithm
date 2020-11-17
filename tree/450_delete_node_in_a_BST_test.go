package tree

// 450. Delete Node in a BST
// 先判断等于和先判断小于有性能差异，是否是分支预测导致？
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		newRoot := root.Right
		for newRoot.Left != nil {
			newRoot = newRoot.Left
		}
		root.Val = newRoot.Val
		root.Right = deleteNode(root.Right, newRoot.Val)
	}
	return root
}
