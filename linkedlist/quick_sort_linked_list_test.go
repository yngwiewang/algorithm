package linkedlist

import (
	"fmt"
	"testing"
)

// leetcode 148

//func partitionSort(head, tail *ListNode) *ListNode{
//	if head == tail {
//		return head
//	}
//
//	pivot, head := head, head.Next
//	small, big := &ListNode{0, nil}, &ListNode{0, nil}
//	preSmall, preBig := &ListNode{0, small}, &ListNode{0, big}
//	var moveSmall, moveBig bool
//	for ; head != tail.Next; head = head.Next {
//		if head.Val >= pivot.Val {
//			big.Next = head
//			big = big.Next
//			moveBig = true
//		} else {
//			small.Next = head
//			small = small.Next
//			moveSmall = true
//		}
//	}
//	small.Next, pivot.Next, big.Next = nil, nil, nil
//	if moveSmall {
//		preSmall.Next.Next = partitionSort(preSmall.Next.Next, small)
//		for ;small.Next != nil;small=small.Next{}
//	}
//	if moveBig {
//		preBig.Next.Next = partitionSort(preBig.Next.Next, big)
//	}
//	small.Next, pivot.Next = pivot, preBig.Next.Next
//	return preSmall.Next.Next
//}
//
//func sortList(head *ListNode) *ListNode {
//	if head == nil{
//		return head
//	}
//	pre := &ListNode{0, head}
//	for head.Next != nil {
//		head = head.Next
//	}
//	tail := head
//	res := partitionSort(pre.Next, tail)
//	return res
//}

//  method2

//func partitionSort(head, tail, pre, last *ListNode) {
//	if head == tail {
//		pre.Next, tail.Next = head, last
//		return
//	}
//
//	pivot, head := head, head.Next
//	small, big := &ListNode{0, nil}, &ListNode{0, nil}
//	preSmall, preBig := &ListNode{0, small}, &ListNode{0, big}
//	var moveSmall, moveBig bool
//	for ; head != tail.Next; head = head.Next {
//		if head.Val >= pivot.Val {
//			big.Next = head
//			big = big.Next
//			moveBig = true
//		} else {
//			small.Next = head
//			small = small.Next
//			moveSmall = true
//		}
//	}
//	//small.Next, pivot.Next, big.Next = nil, nil, nil
//	//pre.Next,big.Next = preSmall.Next.Next, last
//	if moveSmall {
//		pre.Next = preSmall.Next.Next
//	} else {
//		pre.Next = pivot
//	}
//	big.Next = last
//	small.Next, pivot.Next = pivot, preBig.Next.Next
//	if moveSmall {
//		partitionSort(preSmall.Next.Next, small, pre, pivot)
//		//for ; small.Next != nil; small = small.Next {
//		//}
//	}
//	if moveBig {
//		partitionSort(preBig.Next.Next, big, pivot, big.Next)
//	}
//	//small.Next, pivot.Next = pivot, preBig.Next.Next
//
//}
//
//func sortList(head *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//	pre := &ListNode{0, head}
//	for head.Next != nil {
//		head = head.Next
//	}
//	tail := head
//	partitionSort(pre.Next, tail, pre, nil)
//	return pre.Next
//}

// 用小于前一个结点的结点作为pivot

//func partitionSort(pre, head, pivot, tail, last *ListNode) {
//	small, big := &ListNode{0, nil}, &ListNode{0, nil}
//	preBig := big
//	pre.Next = small
//	for ; head != tail.Next; head = head.Next {
//		if head == pivot {
//			continue
//		}
//		if head.Val < pivot.Val {
//			small.Next = head
//			small = small.Next
//		} else {
//			big.Next = head
//			big = big.Next
//		}
//	}
//	small.Next, pivot.Next, big.Next = pivot, preBig.Next, last
//	pre.Next = pre.Next.Next
//	for h := pre.Next; h != pivot; h = h.Next {
//		if h.Next.Val < h.Val {
//			p := h.Next
//			partitionSort(pre, pre.Next, p, small, pivot)
//		}
//	}
//	h := pivot
//	for ; h.Val == h.Next.Val; h = h.Next {
//	}
//
//	for n := h.Next; n != last && n.Next != nil; n = n.Next {
//		if n.Next.Val < n.Val {
//			p := n.Next
//			partitionSort(h, h.Next, p, big, last)
//		}
//	}
//}
//
//func sortList(head *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//	pre := &ListNode{0, head}
//	flag := false
//	for ; head.Next != nil; head = head.Next {
//		if head.Next.Val < head.Val {
//			flag = true
//			break
//		}
//	}
//	if flag == false {
//		return pre.Next
//	}
//	p := head.Next
//	for ; head.Next != nil; head = head.Next {
//	}
//	partitionSort(pre, pre.Next, p, head, nil)
//
//	return pre.Next
//}

// 用头结点作为pivot
//func partitionSort(pre, head, tail *ListNode) {
//	if head == tail {
//		return
//	}
//	small, big := &ListNode{0, nil}, &ListNode{0, nil}
//	preSmall, preBig := small, big
//	pivot := head
//	moveSmall, moveBig := false, false
//	for node := head.Next; node != tail; node = node.Next {
//		if node.Val < pivot.Val {
//			small.Next = node
//			small = small.Next
//			moveSmall = true
//		} else {
//			big.Next = node
//			big = big.Next
//			moveBig = true
//		}
//	}
//	if !moveSmall && !moveBig{
//		return
//	}
//	if moveSmall {
//		small.Next = head
//		pre.Next = preSmall.Next
//		partitionSort(pre, pre.Next, pivot)
//	} else {
//		pre.Next = head
//	}
//	if moveBig {
//		head.Next, big.Next = preBig.Next, tail
//		partitionSort(pivot, pivot.Next, tail)
//	} else {
//		head.Next = tail
//	}
//}
//
//func sortList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	pre := &ListNode{0, head}
//	partitionSort(pre, head, nil)
//	return pre.Next
//}

func TestPartitionSort1(t *testing.T) {
	r := &ListNode{
		Val: 4,
		Next: &ListNode{
			4,
			&ListNode{
				1,
				&ListNode{
					4,
					&ListNode{
						5,
						nil,
					},
				},
			},
		},
	}

	l := sortList(r)
	for ; l != nil; l = l.Next {
		fmt.Printf("%d ", l.Val)

	}
}

//func TestQuickSortLinkedList(t *testing.T) {
//	s1 := &ListNode{
//		Val: 9,
//		Next: &ListNode{
//			Val: 2,
//			Next: &ListNode{
//				Val: 1,
//				Next: &ListNode{
//					Val:  4,
//					Next: nil,
//				},
//			},
//		},
//	}
//	s2 := &ListNode{
//		Val:  2,
//		Next: &ListNode{1, nil}}
//	s3 := &ListNode{
//		Val: 1,
//		Next: &ListNode{
//			Val: 1,
//			Next: &ListNode{
//				Val: 5,
//				Next: &ListNode{
//					Val:  4,
//					Next: nil,
//				},
//			},
//		},
//	}
//	s4 := &ListNode{
//		Val: 2,
//		Next: &ListNode{
//			Val: 2,
//			Next: &ListNode{
//				Val: 1,
//				Next: &ListNode{
//					Val:  4,
//					Next: nil,
//				},
//			},
//		},
//	}
//	s5 := &ListNode{
//		Val: 1,
//		Next: &ListNode{
//			Val: 2,
//			Next: &ListNode{
//				Val: 3,
//				Next: &ListNode{
//					Val:  -4,
//					Next: nil,
//				},
//			},
//		},
//	}
//	l1 := sortList(s1)
//	l2 := sortList(s2)
//	l3 := sortList(s3)
//	l4 := sortList(s4)
//	l5 := sortList(s5)
//	for l1 != nil {
//		fmt.Printf("%d ", l1.Val)
//		l1 = l1.Next
//	}
//	fmt.Println()
//	for l2 != nil {
//		fmt.Printf("%d ", l2.Val)
//		l2 = l2.Next
//	}
//	fmt.Println()
//	for l3 != nil {
//		fmt.Printf("%d ", l3.Val)
//		l3 = l3.Next
//	}
//	fmt.Println()
//	for l4 != nil {
//		fmt.Printf("%d ", l4.Val)
//		l4 = l4.Next
//	}
//	fmt.Println()
//	for l5 != nil {
//		fmt.Printf("%d ", l5.Val)
//		l5 = l5.Next
//	}
//}
