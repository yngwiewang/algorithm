package tree

import (
	"strconv"
	"testing"
)

// 257. Binary Tree Paths
// backtracking
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}
	btpHelper(&res, "", root)
	return res
}

func btpHelper(res *[]string, tmp string, cur *TreeNode) {
	tmp += strconv.Itoa(cur.Val)

	if cur.Left == nil && cur.Right == nil {
		*res = append(*res, tmp)
		return
	}
	tmp += "->"
	if cur.Left != nil {
		btpHelper(res, tmp, cur.Left)
	}
	if cur.Right != nil {
		btpHelper(res, tmp, cur.Right)
	}
}

func Test_btp(t *testing.T) {
	a := arrayToBinaryTree([]int{1, 2, 3, 5})
	res := binaryTreePaths(a)
	t.Log(res)
}
