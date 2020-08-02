package linkedlist

import (
	"fmt"
	"testing"
)

// leetcode 148

//func partitionSort(head, tail *ListNode) *ListNode {
//	if head == tail {
//		return head
//	}
//	pivot, head := head, head.Next
//	small, big := &ListNode{0, nil}, &ListNode{0, nil}
//	preSmall, preBig := small, big
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
//		preSmall.Next = partitionSort(preSmall.Next, small)
//		for ; small.Next != nil; small = small.Next {
//		}
//	}
//	if moveBig {
//		preBig.Next = partitionSort(preBig.Next, big)
//	}
//	small.Next, pivot.Next = pivot, preBig.Next
//	return preSmall.Next
//}
//
//func sortList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
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
//	pivot, head := head, head.Next
//	small, big := &ListNode{0, nil}, &ListNode{0, nil}
//	preSmall, preBig := small, big
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
//	small.Next = pivot
//	if moveSmall {
//		pre.Next = preSmall.Next
//	} else {
//		pre.Next = pivot
//	}
//	if moveBig{
//		pivot.Next,big.Next = preBig.Next, last
//	} else{
//		pivot.Next = last
//	}
//	if moveSmall {
//		partitionSort(preSmall.Next, small, pre, pivot)
//	}
//	if moveBig {
//		partitionSort(preBig.Next, big, pivot, big.Next)
//	}
//
//}
//
//func sortList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
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

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pivot := head.Val
	small, mid, big := new(ListNode), new(ListNode), new(ListNode)
	preSmall, preMid, preBig := small, mid, big
	for ; head != nil; head = head.Next {
		if head.Val < pivot {
			small.Next = head
			small = head
		} else if head.Val > pivot {
			big.Next = head
			big = head
		} else if head.Val == pivot {
			mid.Next = head
			mid = head
		}
	}
	small.Next, mid.Next, big.Next = nil, nil, nil
	return concat(sortList(preSmall.Next), preMid.Next, sortList(preBig.Next))
}

func concat(l, m, h *ListNode) *ListNode {
	pre := new(ListNode)
	if l != nil {
		pre.Next = l
		for ; l.Next != nil; l = l.Next {
		}
		l.Next = m
	} else {
		pre.Next = m
	}
	if h != nil {
		for ; m.Next != nil; m = m.Next {
		}
		m.Next = h
	}

	return pre.Next
}

func TestPartitionSort1(t *testing.T) {
	r := arrayToLinkedList([]int{5, 1, 4, 3, 1, 6, 2})

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

func sortListQuick20200718(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head.Val
	l, m, r := new(ListNode), new(ListNode), new(ListNode)
	prel, prem, prer := l, m, r
	for ; head != nil; head = head.Next {
		if head.Val < p {
			l.Next = head
			l = head
		} else if head.Val > p {
			r.Next = head
			r = head
		} else {
			m.Next = head
			m = head
		}
	}
	l.Next, r.Next = nil, nil
	prel.Next = sortListQuick20200718(prel.Next)
	prer.Next = sortListQuick20200718(prer.Next)
	tmp := prel
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next, m.Next = prem.Next, prer.Next
	return prel.Next
}

func TestSortList20200718(t *testing.T) {
	a := arrayToLinkedList([]int{7})
	l := sortListQuick20200718(a)
	fmt.Println(l)
}
