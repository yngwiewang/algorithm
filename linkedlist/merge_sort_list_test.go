package linkedlist

import (
	"fmt"
	"testing"
)

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

func sortListMerge20200718(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	pre := &ListNode{}
	left := head
	right := slow.Next
	slow.Next = nil
	pre.Next = merge20200718(sortListMerge20200718(left), sortListMerge20200718(right))
	return pre.Next
}

func merge20200718(l, r *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	for l != nil && r != nil {
		if l.Val < r.Val {
			cur.Next = l
			cur = l
			l = l.Next
		} else {
			cur.Next = r
			cur = r
			r = r.Next
		}
	}
	if l != nil {
		cur.Next = l
	}
	if r != nil {
		cur.Next = r
	}
	return pre.Next
}

func TestMergeSortLinkedList20200718(t *testing.T) {
	a := arrayToLinkedList([]int{5, 4, 3, 2, 1})
	l := sortListMerge20200718(a)
	for ; l != nil; l = l.Next {
		fmt.Printf("%d ", l.Val)
	}
	fmt.Println()
}
