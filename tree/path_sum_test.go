package tree

func hasPathSum(root *TreeNode, sum int) bool {
	var res bool
	if root == nil {
		return false
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			res = true
		}
		return false
	}
	if res == true {
		return true
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

func hasPathSum1(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			return true
		}
		return false
	}
	if hasPathSum(root.Left, sum) {
		return true
	}
	return hasPathSum(root.Right, sum)
}

func helper(root *TreeNode, sum *int) bool {
	if root == nil {
		return false
	}
	*sum -= root.Val
	if root.Left == nil && root.Right == nil {
		if *sum == 0 {
			return true
		}
		return false
	}
	if hasPathSum(root.Left, *sum) {
		return true
	}
	return hasPathSum(root.Right, *sum)
}

func hasPathSum2(root *TreeNode, sum int) bool {
	return helper(root, &sum)
}

func hasPathSum3(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
