package tree

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	head := root
	insertIntoBSTHelper(head, val)
	return root
}

func insertIntoBSTHelper(root *TreeNode, val int) {
	if val > root.Val {
		if root.Right == nil {
			root.Right = &TreeNode{val, nil, nil}
			return
		} else {
			insertIntoBSTHelper(root.Right, val)
		}
	} else {
		if root.Left == nil {
			root.Left = &TreeNode{val, nil, nil}
			return
		} else {
			insertIntoBSTHelper(root.Left, val)
		}
	}
}

func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	head := root
	if root == nil {
		return &TreeNode{val, nil, nil}
	}

	for {
		if val > root.Val {
			if root.Right == nil {
				root.Right = &TreeNode{val, nil, nil}
				break
			} else {
				root = root.Right
			}
		} else {
			if root.Left == nil {
				root.Left = &TreeNode{val, nil, nil}
				break
			} else {
				root = root.Left
			}
		}
	}
	return head
}

func insertIntoBST2(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return &TreeNode{val, nil, nil}
	}

	if val > root.Val {
		root.Right = insertIntoBST2(root.Right, val)
	} else {
		root.Left = insertIntoBST2(root.Left, val)
	}
	return root
}
