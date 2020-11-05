package tree

// 116. Populating Next Right Pointers in Each Node116

// Definition for a Node116.
 type Node116 struct {
     Val int
     Left *Node116
     Right *Node116
     Next *Node116
 }


// recursive two Node116s with helper, slow
func connect(root *Node116) *Node116 {
	if root != nil {
		connectHelper(root.Left, root.Right)
	}
	return root
}

func connectHelper(l, r *Node116) {
	if l == nil || r == nil {
		return
	}
	l.Next = r
	if l.Right != nil {
		l.Right.Next = r.Left
		connectHelper(l.Right, r.Left)
	}
	connectHelper(l.Left, l.Right)
	connectHelper(r.Left, r.Right)
}

// level iterate with queue
func connectA(root *Node116) *Node116 {
	if root == nil {
		return nil
	}
	q := []*Node116{root}
	for len(q) > 0 {
		length := len(q)
		for i := range q {
			tmp := q[0]
			q = q[1:]
			if i != length-1 {
				tmp.Next = q[0]
			}
			if tmp.Left != nil {
				q = append(q, tmp.Left)
				q = append(q, tmp.Right)
			}
		}
	}
	return root
}

// recursive single Node116, faster
func connectB(root *Node116) *Node116 {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		root.Left.Next = root.Right
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
	}
	connectB(root.Left)
	connectB(root.Right)
	return root
}

//  a little varaition of B
func connectC(root *Node116) *Node116 {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	}
	root.Left.Next = root.Right
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}

	connectC(root.Left)
	connectC(root.Right)
	return root
}
