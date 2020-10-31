package stack

import (
	"reflect"
	"testing"
)

// 496. Next Greater Element I

// 单调栈 100%
func nextGreaterElementA(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	m := make(map[int]int, len(nums1))
	stack := []int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums2[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			m[nums2[i]] = -1
		} else {
			m[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	for i, v := range nums1 {
		res[i] = m[v]
	}
	return res
}

// 在 nums2 中遍历寻找 100%
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	m := make(map[int]int, len(nums2))
	for i, v := range nums2 {
		m[v] = i
	}
	var found bool
	for i, v := range nums1 {
		found = false
		for j := m[v] + 1; j < len(nums2); j++ {
			if nums2[j] > v {
				res[i] = nums2[j]
				found = true
				break
			}
		}
		if !found {
			res[i] = -1
		}
	}
	return res
}

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{4, 1, 2}, []int{1, 3, 4, 2}}, []int{-1, 3, -1}},
		{"2", args{[]int{4, 1, 2}, []int{1, 2, 3, 4}}, []int{-1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
