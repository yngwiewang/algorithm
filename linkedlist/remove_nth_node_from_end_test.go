package linkedlist

import (
	"reflect"
	"testing"
)

// 19.Remove Nth Node From End of List

// 自己的解法，边界判断比较复杂，容易出错
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	i, j := 0, 0
	pre := head
	for h := head; h.Next != nil; h = h.Next {
		if i >= n {
			pre = pre.Next
			j++
		}
		i++
	}
	if n == i {
		head.Next = head.Next.Next
		return head
	}
	if j == 0 {
		return head.Next
	}
	pre.Next = pre.Next.Next
	return head
}

// * leetcode最快解法，简洁干净，链表题还是要考虑用dummy节点
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	pre := &ListNode{0, head}
	slow, fast := pre, pre
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast, slow = fast.Next, slow.Next
	}
	slow.Next = slow.Next.Next

	return pre.Next
}

// ! leetcode 内存最小答案，比上面的方法多了一步判断，减少了一个节点指针的内存消耗
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	pre := &ListNode{0, head}
	fast := pre
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil{
		return head.Next
	}
	slow := pre
	for fast.Next != nil {
		fast, slow = fast.Next, slow.Next
	}
	slow.Next = slow.Next.Next

	return pre.Next
}

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"", args{arrayToLinkedList([]int{1}), 1}, arrayToLinkedList([]int{})},
		{"", args{arrayToLinkedList([]int{1, 2}), 1}, arrayToLinkedList([]int{1})},
		{"", args{arrayToLinkedList([]int{1, 2}), 2}, arrayToLinkedList([]int{2})},
		{"", args{arrayToLinkedList([]int{1, 2, 3}), 1}, arrayToLinkedList([]int{1, 2})},
		{"", args{arrayToLinkedList([]int{1, 2, 3}), 2}, arrayToLinkedList([]int{1, 3})},
		{"", args{arrayToLinkedList([]int{1, 2, 3}), 3}, arrayToLinkedList([]int{2, 3})},
		{"", args{arrayToLinkedList([]int{1, 2, 3, 4}), 1}, arrayToLinkedList([]int{1, 2, 3})},
		{"", args{arrayToLinkedList([]int{1, 2, 3, 4}), 2}, arrayToLinkedList([]int{1, 2, 4})},
		{"", args{arrayToLinkedList([]int{1, 2, 3, 4}), 3}, arrayToLinkedList([]int{1, 3, 4})},
		{"", args{arrayToLinkedList([]int{1, 2, 3, 4}), 4}, arrayToLinkedList([]int{2, 3, 4})},
		{"", args{arrayToLinkedList([]int{1, 2, 3, 4, 5}), 1}, arrayToLinkedList([]int{1, 2, 3, 4})},
		{"", args{arrayToLinkedList([]int{1, 2, 3, 4, 5}), 2}, arrayToLinkedList([]int{1, 2, 3, 5})},
		{"", args{arrayToLinkedList([]int{1, 2, 3, 4, 5}), 3}, arrayToLinkedList([]int{1, 2, 4, 5})},
		{"", args{arrayToLinkedList([]int{1, 2, 3, 4, 5}), 4}, arrayToLinkedList([]int{1, 3, 4, 5})},
		{"", args{arrayToLinkedList([]int{1, 2, 3, 4, 5}), 5}, arrayToLinkedList([]int{2, 3, 4, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd2(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
