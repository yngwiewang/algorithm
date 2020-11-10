package tree

import "math"

// 654. Maximum Binary Tree

// recursive
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxV := math.MinInt32
	maxP := 0
	for i, v := range nums {
		if v > maxV {
			maxV = v
			maxP = i
		}
	}
	root := &TreeNode{maxV, constructMaximumBinaryTree(nums[:maxP]),
		constructMaximumBinaryTree(nums[maxP+1:])}

	return root
}
