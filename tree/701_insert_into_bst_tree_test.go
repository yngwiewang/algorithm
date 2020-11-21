package tree

import "testing"

// 701. Insert into a Binary Search Tree
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

// review 20201121
func insertIntoBST202011211(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{val, nil, nil}
	if root == nil {
		return newNode
	}
	n := root
	for n.Left != nil || n.Right != nil {
		if val < n.Val {
			if n.Left == nil {
				n.Left = newNode
				return root
			}
			n = n.Left
		} else {
			if n.Right == nil {
				n.Right = newNode
				return root
			}
			n = n.Right
		}
	}
	if val < n.Val {
		n.Left = newNode
	} else {
		n.Right = newNode
	}
	return root
}

func insertIntoBST202011212(root *TreeNode, val int) *TreeNode {
	if root == nil{
		return &TreeNode{val, nil, nil}
	}
	if val < root.Val{
		root.Left = insertIntoBST202011212(root.Left, val)
	} else {
		root.Right = insertIntoBST202011212(root.Right, val)
	}
	return root
}
func Test_insertBST20201121(t *testing.T) {
	c := Constructor()
	tree := c.deserialize("8,#,55,39,#,11,#,#,23,#,#,")
	instered := insertIntoBST202011212(tree, 17)
	t.Log(inorderTraversal(instered))
}
