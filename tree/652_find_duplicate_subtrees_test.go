package tree

import (
	"fmt"
	"strconv"
	"strings"
)

// 652. Find Duplicate Subtrees
// preorder traversal
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res := make([]*TreeNode, 0)
	m := make(map[string]int, 0)
	findDupSubtrees(root, &res, m)
	return res
}

func findDupSubtrees(root *TreeNode, res *[]*TreeNode, m map[string]int) string {
	if root == nil {
		return "#"
	}
	subtree := strconv.Itoa(root.Val) + "," + findDupSubtrees(root.Left, res, m) + "," + findDupSubtrees(root.Right, res, m)
	if m[subtree] == 1 {
		*res = append(*res, root)
	}
	m[subtree]++
	return subtree
}

// use strings.Builder, faster
func findDupSubtrees1(root *TreeNode, res *[]*TreeNode, m map[string]int) string {
	if root == nil {
		return "#"
	}
	var sb strings.Builder
	fmt.Fprintf(&sb, "%d,%s%s", root.Val, findDupSubtrees(root.Left, res, m), findDupSubtrees(root.Right, res, m))
	subtree := sb.String()
	if m[subtree] == 1 {
		*res = append(*res, root)
	}
	m[subtree]++
	return subtree
}
