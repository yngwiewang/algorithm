package linkedlist

// 138. Copy List with Random Pointer
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 复制节点跟在原节点后面,1-1'-2-2'-3-3'
// 分离出1-2-3和1‘'-2'-3'
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	for h := head; h != nil; h = h.Next {
		tmp := &Node{h.Val, h.Next, h.Random}
		h.Next = tmp
		h = tmp
	}
	for h := head; h != nil; h = h.Next.Next {
		if h.Random != nil {
			h.Next.Random = h.Random.Next
		}
	}
	cur, clonedHead := head, head.Next
	for cur != nil && cur.Next != nil {
		tmp := cur.Next
		cur.Next = tmp.Next
		cur = tmp
	}
	return clonedHead
}

func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}

	var newHead *Node
	i := 0
	// map (old node -> new node)
	nodeMap := map[*Node]*Node{}

	for head != nil {
		n, ok := nodeMap[head]
		if !ok {
			n = &Node{Val: head.Val}
			nodeMap[head] = n
		}

		if _, ok := nodeMap[head.Next]; !ok && head.Next != nil {
			nodeMap[head.Next] = &Node{Val: head.Next.Val}
		}
		n.Next = nodeMap[head.Next]

		if _, ok := nodeMap[head.Random]; !ok && head.Random != nil {
			nodeMap[head.Random] = &Node{Val: head.Random.Val}
		}
		n.Random = nodeMap[head.Random]

		if i == 0 {
			newHead = n
		}
		i++
		head = head.Next
	}
	return newHead
}
