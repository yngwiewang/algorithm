package tree

import "container/list"

// 111. Minimum Depth of Binary Tree
// BFS, use slice
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		res++
		// fmt.Println(queue[0].Val)
		for i := 0; i < size; i++ {
			// node := queue[0]
			// queue = queue[1:]
			node := queue[i]
			if node.Left == nil && node.Right == nil {
				return res
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// ! 一次弹出，提高性能
		queue = queue[size:]
	}
	return res
}

// BFS, use linkedlist
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		res++
		// fmt.Println(queue[0].Val)
		for i := 0; i < size; i++ {
			elem := queue.Front()
			queue.Remove(elem)
			node := elem.Value.(*TreeNode)
			if node.Left == nil && node.Right == nil {
				return res
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return res
}
