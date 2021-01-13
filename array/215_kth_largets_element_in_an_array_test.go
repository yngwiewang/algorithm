package array

import (
	"sort"
	"testing"
)

// 215. Kth Largest Element in an Array

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 2, 1, 6, 5, 4}, 2}, 5},
		{"2", args{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4}, 4},
		{"3", args{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 6}, 3},
		{"4", args{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 1}, 6},
		{"5", args{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 9}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
