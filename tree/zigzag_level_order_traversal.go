package tree

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	nodes := []*TreeNode{root}
	order := 0
	for len(nodes) > 0 {
		var tmp []int
		var tmpNodes []*TreeNode
		if order == 0 {
			for _, node := range nodes {
				tmp = append(tmp, node.Val)
				if node.Left != nil {
					tmpNodes = append(tmpNodes, node.Left)
				}
				if node.Right != nil {
					tmpNodes = append(tmpNodes, node.Right)
				}
			}
			order = 1
		} else {
			for i := len(nodes) - 1; i >= 0; i-- {
				tmp = append(tmp, nodes[i].Val)

				if nodes[i].Right != nil {
					tmpNodes = append([]*TreeNode{nodes[i].Right}, tmpNodes...)
				}
				if nodes[i].Left != nil {
					tmpNodes = append([]*TreeNode{nodes[i].Left}, tmpNodes...)
				}
			}
			order = 0
		}
		nodes = tmpNodes
		res = append(res, tmp)
	}
	return res
}
