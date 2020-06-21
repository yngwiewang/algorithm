package linkedlist

import (
	"fmt"
	"testing"
)

// leetcode 2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		pSum  ListNode
		p     *ListNode
		pTail = &pSum
		value int
		carry int
	)

	for l1 != nil && l2 != nil {
		value = l1.Val + l2.Val + carry
		carry = value / 10
		value = value % 10
		pCur := ListNode{
			value,
			nil,
		}
		pTail.Next = &pCur
		pTail = &pCur
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil {
		p = l1
	} else if l2 != nil {
		p = l2
	}

	for p != nil {
		value = p.Val + carry
		carry = value / 10
		value %= 10
		pCur := ListNode{
			value,
			nil,
		}
		pTail.Next = &pCur
		pTail = &pCur
		p = p.Next
	}

	if carry > 0 {
		pTail.Next = &ListNode{carry, nil}
	}

	return pSum.Next
}

func AddTwoNumbers1(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	sum := 0

	for l1 != nil || l2 != nil || sum > 0 {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		p.Next = &ListNode{sum % 10, nil}
		p = p.Next
		sum = sum / 10
	}
	return head.Next
}

func TestAddTwoNumbers(t *testing.T) {
	c := ListNode{
		Val:  9,
		Next: nil,
	}
	b := ListNode{
		Val:  9,
		Next: &c,
	}
	l1 := &ListNode{
		Val:  2,
		Next: &b,
	}
	d := ListNode{
		Val:  9,
		Next: nil,
	}
	e := ListNode{
		Val:  9,
		Next: &d,
	}
	//var l1 = new(ListNode)
	//var l2 = new(ListNode)

	l2 := &ListNode{
		Val:  2,
		Next: &e,
	}
	//res := addTwoNumbers(l1, l2)
	//for res != nil {
	//	fmt.Println(res.Val)
	//	res = res.Next
	//}
	res := AddTwoNumbers1(l1, l2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
