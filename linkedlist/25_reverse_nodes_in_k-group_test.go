package linkedlist

import (
	"testing"
)

//  25. Reverse Nodes in k-Group
// 递归，提取函数反转A、B两点间（补包含B）的节点
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head // 如果剩余元素不足K个就不反转
		}
		b = b.Next
	}
	newHead := reverseAtoB(a, b)
	a.Next = reverseKGroup(b, k)

	return newHead
}

func reverseAtoB(a, b *ListNode) *ListNode {
	if a == nil {
		return nil
	}
	var pre *ListNode
	for a != b {
		tmp := a
		a = a.Next
		tmp.Next, pre = pre, tmp
	}
	return pre
}

func Test_reverseKGroup(t *testing.T) {
	a := arrayToLinkedList([]int{1, 2, 3, 4})
	res := reverseKGroup(a, 2)
	t.Log(res)
}
