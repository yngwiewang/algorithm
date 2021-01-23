package linkedlist

import (
	"fmt"
	"testing"
)

// 725. Split Linked List in Parts

func splitListToParts(root *ListNode, k int) []*ListNode {
	res := make([]*ListNode, k)
	length := 0
	for head := root; head != nil; head = head.Next {
		length++
	}
	if length <= k {
		for i, head := 0, root; i < length; i++ {
			res[i] = head
			head = head.Next
		}
		for i := 0; i < length; i++ {
			res[i].Next = nil
		}
		return res
	}
	remainder, quotient := length/k, length%k

	var r *ListNode
	i := 0
	for ; i < quotient; i++ {
		root, r = splitListToTwoParts(root, remainder+1)
		res[i], root = root, r
	}
	for j := i; j < k; j++ {
		root, r = splitListToTwoParts(root, remainder)
		res[j], root = root, r
	}

	return res
}

func splitListToTwoParts(root *ListNode, k int) (left, right *ListNode) {
	left, tmp := root, root
	for i := 0; i < k-1; i++ {
		tmp = tmp.Next
	}
	right = tmp.Next
	tmp.Next = nil
	return
}

func TestSplitListToTwoParts(t *testing.T) {
	a := arrayToLinkedList([]int{1, 2, 3, 4, 5})
	l, r := splitListToTwoParts(a, 1)
	fmt.Println(l)
	fmt.Println(r)
}

func TestSplitListToParts(t *testing.T) {
	a := arrayToLinkedList([]int{1, 2, 3, 4, 5})
	res := splitListToParts(a, 5)

	fmt.Println(len(res))
	for i, v := range res {
		fmt.Println(i, v)
	}
}
