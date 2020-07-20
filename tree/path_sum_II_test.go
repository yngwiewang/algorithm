package tree

func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var cur []int
	helperPathSum(root, sum, cur, &res)
	return res
}

func helperPathSum(root *TreeNode, sum int, cur []int, res *[][]int) {
	sum -= root.Val
	cur = append(cur, root.Val)
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			*res = append(*res, tmp)
			return
		}
	}
	if root.Left != nil {
		helperPathSum(root.Left, sum, cur, res)
	}
	if root.Right != nil {
		helperPathSum(root.Right, sum, cur, res)
	}
	cur = cur[:len(cur)-1]
}
