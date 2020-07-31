package linkedlist

import "testing"

// 使用额外的slice存储全部节点
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	var tmp []*ListNode
	for node := head; node != nil; node = node.Next {
		tmp = append(tmp, node)
	}
	for i := range tmp {
		tmp[i].Next = nil
	}
	length := len(tmp)
	node := head
	for i := 1; i <= length/2; i++ {
		node.Next = tmp[length-i]
		node = node.Next
		if i == length/2 && length%2 == 0 {
			break
		}
		node.Next = tmp[i]
		node = node.Next
	}
	return
}

func TestReorderList(t *testing.T) {
	a := arrayToLinkedList([]int{1, 2, 3, 4})
	reorderList2(a)
	t.Log(a)
}

// same speed as method1, a little less mem
func reorderList1(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	var tmp []*ListNode
	for node := head.Next; node != nil; node = node.Next {
		tmp = append(tmp, node)
	}
	for i := range tmp {
		tmp[i].Next = nil
	}
	node := head
	for len(tmp) > 0 {
		node.Next = tmp[len(tmp)-1]
		tmp = tmp[:len(tmp)-1]
		node = node.Next
		if len(tmp) > 0 {
			node.Next = tmp[0]
			tmp = tmp[1:]
			node = node.Next
		}
	}
	return
}

// 从中点切分链表，后半部分反转，然后逐个合并
func reorderList2(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	l := head
	r := slow.Next
	slow.Next = nil
	rr := reverseList(r)
	for rr != nil {
		tmp := rr
		rr = rr.Next
		tmp.Next = l.Next
		l.Next = tmp
		l = l.Next.Next
	}
}
