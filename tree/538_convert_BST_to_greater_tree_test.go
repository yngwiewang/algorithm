package tree

import "testing"

// 538. Convert BST to Greater Tree
// postorder traversal
// ! 把节点为空的判断放在顶层，Helper内判断子树不为空才递归
// ! 可以显著提高性能
func convertBST1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	sum := 0
	convertBSTHelper1(root, &sum)
	return root
}

func convertBSTHelper1(root *TreeNode, sum *int) {
	if root.Right != nil {
		convertBSTHelper1(root.Right, sum)
	}
	*sum += root.Val
	root.Val = *sum
	if root.Left != nil {
		convertBSTHelper1(root.Left, sum)
	}
}

// use stack, stupid
func convertBST2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	bstToStack(root, &stack)
	sum := 0
	for i := len(stack) - 1; i >= 0; i-- {
		sum += stack[i].Val
		stack[i].Val = sum
	}
	return root
}

func bstToStack(root *TreeNode, stack *[]*TreeNode) {
	if root.Left != nil {
		bstToStack(root.Left, stack)
	}
	*stack = append(*stack, root)
	if root.Right != nil {
		bstToStack(root.Right, stack)
	}
}

// use stack, smarter
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	stack := make([]*TreeNode, 0)
	node := root
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Right
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sum += root.Val
		root.Val = sum
		root = root.Left

	}
	return node
}

func Test_convertBST(t *testing.T) {
	c := Constructor()
	tree := c.deserialize("4,1,6,0,2,5,7,#,#,#,3,#,#,#,8,#,#,#,#,")
	converted := convertBST(tree)
	treeS := c.serialize(converted)
	t.Log(treeS)
}
