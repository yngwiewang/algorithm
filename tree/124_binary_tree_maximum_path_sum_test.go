package tree

import (
	"math"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 124. Binary Tree Maximum Path Sum

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	if root == nil {
		return res
	}
	helperMaxPathSum(root, &res)
	return res
}

func helperMaxPathSum(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	l := helperMaxPathSum(root.Left, res)
	r := helperMaxPathSum(root.Right, res)
	tmp := max(max(l, r)+root.Val, root.Val)
	ans := max(tmp, l+r+root.Val)
	*res = max(ans, *res)
	return tmp
}

func maxMulti(a ...int) int {
	m := a[0]
	for _, i := range a {
		if i > m {
			m = i
		}
	}
	return m
}

func TestMaxMulti(t *testing.T) {
	r := maxMulti(3, 7, 5, 2, 4, 6, 1, 9)
	t.Log(r)
}

func maxPathSum1(root *TreeNode) int {
	res := math.MinInt32
	if root == nil {
		return res
	}
	helperMaxPathSum1(root, &res)
	return res
}

func helperMaxPathSum1(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	var left, right int
	if l := helperMaxPathSum1(root.Left, res); l > 0 {
		left = l
	}
	if r := helperMaxPathSum1(root.Right, res); r > 0 {
		right = r
	}
	if left+right+root.Val > *res {
		*res = left + right + root.Val
	}
	if left > right {
		return left + root.Val
	}
	return right + root.Val
}

func maxPathSum2(root *TreeNode) int {
	res := math.MinInt32
	if root == nil {
		return res
	}
	helperMaxPathSum2(root, &res)
	return res
}

func helperMaxPathSum2(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := max(0, helperMaxPathSum2(root.Left, res))
	right := max(0, helperMaxPathSum2(root.Right, res))
	*res = max(*res, left+right+root.Val)
	return max(left, right) + root.Val
}

// review 20200127
func maxPathSumA(root *TreeNode) int {
	var ans int = math.MinInt32
	oneSidePath(root, &ans)
	return ans
}

func oneSidePath(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	left := common.MaxInt(0, oneSidePath(root.Left, ans))
	right := common.MaxInt(0, oneSidePath(root.Right, ans))
	*ans = common.MaxInt(*ans, left+right+root.Val)
	return common.MaxInt(left, right) + root.Val
}
