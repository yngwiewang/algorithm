package linkedlist

import (
	"math"
)

// 1290. Convert Binary Number in a Linked List to Integer

func getDecimalValue(head *ListNode) int {
	a := []int{}
	for head != nil {
		a = append(a, head.Val)
		head = head.Next
	}
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
	res := 0
	for i, v := range a {
		res += v * int(math.Pow(2, float64(i)))
	}
	return res
}

func getDecimalValue1(head *ListNode) int {
	a := []int{}
	for head != nil {
		a = append(a, head.Val)
		head = head.Next
	}

	res := 0
	for i, v := range a {
		res += v * int(math.Pow(2, float64(len(a)-i-1)))
	}
	return res
}
