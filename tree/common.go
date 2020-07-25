package tree

func arrayToBinaryTree(a []int) (root *TreeNode) {
	var nodes []*TreeNode
	for _, v := range a {
		nodes = append(nodes, &TreeNode{v, nil, nil})
	}
	if len(nodes) > 0 {
		root = nodes[0]
	} else {
		return root
	}
	for i := 0; i < len(nodes); i++ {
		if len(nodes) > i*2+1 {
			nodes[i].Left = nodes[i*2+1]
		}
		if len(nodes) > i*2+2 {
			nodes[i].Right = nodes[i*2+2]
		}

	}
	return root
}
