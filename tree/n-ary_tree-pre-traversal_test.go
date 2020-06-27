package tree

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

// recursive
func preorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	for _, node := range root.Children {
		res = append(res, preorder(node)...)
	}
	return res
}

//stack
func preorder1(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		if cur == nil { // 这一步判断可以提高性能
			continue
		}
		res = append(res, cur.Val)
		stack = stack[:len(stack)-1]
		for i := len(cur.Children) - 1; i > -1; i-- {
			stack = append(stack, cur.Children[i])
		}
	}
	return res
}
