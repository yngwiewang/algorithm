package tree

import "testing"

// 700. Search in a Binary Search Tree

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)

}

func Test_searchBST(t *testing.T) {
	c := Constructor()
	tree := c.deserialize("4,2,7,1,3,#,#,#,#,#,#,")
	res := searchBST(tree, 5)
	if res != nil {
		t.Log(res.Val)
	} else {
		t.Log("None")
	}

}
