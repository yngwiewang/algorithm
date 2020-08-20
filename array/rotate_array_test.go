package array

import (
	"fmt"
	"testing"
)

// add 189. Rotate Array
func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	k = k % len(nums)
	new := make([]int, len(nums))
	copy(new[:k], nums[len(nums)-k:])
	copy(new[k:], nums[:len(nums)-k])
	copy(nums, new)
}

func rotate1(nums []int, k int) {
	if k == 0 {
		return
	}
	k = k % len(nums)
	if k <= len(nums)/2 {
		new := make([]int, len(nums)-k)
		copy(new, nums[:len(nums)-k])
		copy(nums[:k], nums[len(nums)-k:])
		copy(nums[k:], new)
	} else {
		new := make([]int, k)
		copy(new, nums[len(nums)-k:])
		copy(nums[k:], nums[:len(nums)-k+1])
		copy(nums, new)
	}
	fmt.Println(nums)
}

func rotate2(nums []int, k int) {
	if k == 0 {
		return
	}
	k = k % len(nums)
	if k <= len(nums)/2 {
		new := make([]int, k)
		copy(new, nums[len(nums)-k:])
		copy(nums[k:], nums[:len(nums)-k])
		copy(nums, new)
	} else {
		new := make([]int, len(nums)-k)
		copy(new, nums[:len(nums)-k])
		copy(nums, nums[len(nums)-k:])
		copy(nums[k:], new)
	}
	fmt.Println(nums)
}

func rotate3(nums []int, k int) {
	if k == 0 {
		return
	}
	k = k % len(nums)
	reverse(nums[:len(nums)-k])
	reverse(nums[len(nums)-k:])
	reverse(nums)
	fmt.Println(nums)
}

func rotate4(nums []int, k int) {
	if k == 0 {
		return
	}
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
	fmt.Println(nums)
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	// fmt.Println(nums)
}

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{"0", args{[]int{1, 2, 3, 4}, 1}},
		{"0", args{[]int{1, 2, 3, 4}, 2}},
		{"0", args{[]int{1, 2, 3, 4, 5, 6, 7}, 2}},
		{"0", args{[]int{1, 2, 3}, 1}},
		{"0", args{[]int{1, 2, 3, 4}, 2}},
		{"0", args{[]int{1, 2, 3, 4}, 3}},
		{"0", args{[]int{1, 2, 3, 4, 5, 6, 7}, 3}},
		{"0", args{[]int{1, 2, 3}, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate4(tt.args.nums, tt.args.k)
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{[]int{1, 2, 3}}},
		{"", args{[]int{1, 2, 3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse(tt.args.nums)
		})
	}
}
