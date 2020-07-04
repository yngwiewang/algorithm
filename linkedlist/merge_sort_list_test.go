package linkedlist

import (
	"fmt"
	"testing"
)

func arrayToLinkedList(a []int) *ListNode {
	l := new(ListNode)
	pre := l
	for _, i := range a {
		l.Next = &ListNode{i, nil}
		l = l.Next
	}
	return pre.Next
}

func merge(l1, l2 *ListNode) *ListNode {
	tmp := new(ListNode)
	pre := tmp
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tmp.Next = l1
			l1 = l1.Next
			tmp = tmp.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
			tmp = tmp.Next
		}
	}
	if l1 != nil {
		tmp.Next = l1
	}
	if l2 != nil {
		tmp.Next = l2
	}

	return pre.Next
}

//func divide(head *ListNode) *ListNode {
//	if head.Next == nil{
//		return head
//	}
//	slow, fast := head, head.Next
//	for fast != nil && fast.Next != nil {
//		slow, fast = slow.Next, fast.Next.Next
//	}
//	mid := slow.Next
//	slow.Next = nil
//	return merge(divide(head), divide(mid))
//}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	return merge(mergeSort(head), mergeSort(mid))
}

func TestMergeSortLinkedList(t *testing.T) {
	a := arrayToLinkedList([]int{9, 2, 7, 4, 5})
	l := mergeSort(a)
	for ; l != nil; l = l.Next {
		fmt.Printf("%d ", l.Val)
	}
	fmt.Println()
}

func TestArrayToLinkedList(t *testing.T) {
	a := []int{1, 3, 5}
	l := arrayToLinkedList(a)
	for ; l != nil; l = l.Next {
		fmt.Printf("%d ", l.Val)
	}
	fmt.Println()
}

func TestMerge(t *testing.T) {
	a := arrayToLinkedList([]int{1, 3, 5})
	b := arrayToLinkedList([]int{2, 4, 6})

	l := merge(a, b)

	for ; l != nil; l = l.Next {
		fmt.Printf("%d ", l.Val)
	}
	fmt.Println()

}
