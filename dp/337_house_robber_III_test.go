package dp

import "github.com/yngwiewang/algorithm/common"

// 337. House Robber III

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func robIII(root *TreeNode) int {
	memo := make(map[*TreeNode]int)
	return robIIIHelper(root, memo)
}

func robIIIHelper(root *TreeNode, memo map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if v, ok := memo[root]; ok {
		return v
	}
	do_cur := root.Val
	if root.Left != nil {
		do_cur += robIIIHelper(root.Left.Left, memo) + robIIIHelper(root.Left.Right, memo)
	}
	if root.Right != nil {
		do_cur += robIIIHelper(root.Right.Left, memo) + robIIIHelper(root.Right.Right, memo)
	}
	not_do_cur := robIIIHelper(root.Left, memo) + robIIIHelper(root.Right, memo)
	memo[root] = common.MaxInt(do_cur, not_do_cur)
	return memo[root]
}
