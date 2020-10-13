package array

import (
	"sort"
	"testing"
)

// 532. K-diff Pairs in an Array
// 排序后逐个判断，跳过重复数字剪枝
func findPairsA(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	res := 0
	i, j := 0, 1
	for i < len(nums) {
		if nums[j]-nums[i] >= k {
			if nums[j]-nums[i] == k {
				res++
			}
			t := i + 1
			for ; t < len(nums) && nums[t] == nums[t-1]; t++ {
			}
			i = t
			if i < len(nums)-1 {
				j = i + 1
			} else {
				break
			}
			continue
		}
		t := j + 1
		for ; t < len(nums) && nums[t] == nums[t-1]; t++ {
		}
		j = t
		if j == len(nums) {
			for ; i < len(nums); i++ {
			}
			if i < len(nums)-1 {
				j = i + 1
			} else {
				break
			}
		}
	}
	return res
}

// use map
func findPairs(nums []int, k int) int {
	m := make(map[int]int)
	res := 0
	for _, v := range nums {
		m[v]++
	}
	for key := range m {
		if (k > 0 && m[key+k] > 0) || (k == 0 && m[key] > 1) {
			res++
		}
	}
	return res
}
func Test_findPairs(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1}, 0}, 0},
		{"-1,-2,-3", args{[]int{-1, -2, -3}, 1}, 2},
		{"1,2,4,4,3,3,0,9,2,3", args{[]int{1, 2, 4, 4, 3, 3, 0, 9, 2, 3}, 3}, 2},
		{"1,3,1,5,4", args{[]int{1, 3, 1, 5, 4}, 0}, 1},
		{"1,2,3,4,5", args{[]int{1, 2, 3, 4, 5}, 1}, 4},
		{"3,1,4,1,5", args{[]int{3, 1, 4, 1, 5}, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPairs(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
