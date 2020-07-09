package tree

// leetcode 559

func maxNaryDepth(root *Node) int {
	var depth int
	if root == nil {
		return depth
	}

	dfs_nary_tree_depth(root, &depth, 0)
	return depth
}

func dfs_nary_tree_depth(node *Node, depth *int, cur int) {
	if node == nil {
		return
	}
	cur++
	if cur > *depth {
		*depth = cur
	}
	for _, n := range node.Children {
		dfs_nary_tree_depth(n, depth, cur)
	}
	cur--
}

func maxNaryDepth1(root *Node) int {
	var depth int
	if root == nil {
		return depth
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		depth++
		length := len(stack)
		for i := 0; i < length; i++ {
			stack = append(stack, stack[i].Children...)
		}
		stack = stack[length:]
	}
	return depth
}

func maxNaryDepth2(root *Node) int {
	var depth int
	if root == nil {
		return depth
	}
	for _, node := range root.Children {
		cur := maxNaryDepth(node)
		if cur > depth {
			depth = cur
		}
	}
	return depth + 1
}
