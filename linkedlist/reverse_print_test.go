package linkedlist

import "testing"

//剑指 Offer 06. 从尾到头打印链表

// 逐个读取链表节点，倒序插入slice
func reversePrint(head *ListNode) []int {
	res := make([]int, 0, 10000)
	if head == nil {
		return res
	}
	for ; head != nil; head = head.Next {
		res = append([]int{head.Val}, res...)
	}
	return res
}

func TestReversePrint(t *testing.T) {
	a := arrayToLinkedList([]int{1, 2, 3, 4, 5})
	b := reversePrint1(a)
	t.Log(b)
}

// 逐个读取链表节点，顺序插入slice，倒序排列slice
func reversePrint1(head *ListNode) []int {
	var res []int
	if head == nil {
		return res
	}
	for ; head != nil; head = head.Next {
		res = append(res, head.Val)
	}
	reverseSlice(res)
	return res
}

func reverseSlice(l []int) {
	for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
		l[i], l[j] = l[j], l[i]
	}
}
