package array

import (
	"sort"
	"testing"
)

// 287. Find the Duplicate Number

// use map
func findDuplicateA(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		if m[v] == true {
			return v
		}
		m[v] = true
	}
	return 0
}

// sort and compare
func findDuplicateB(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return 0
}

// binary search
func findDuplicate(nums []int) int {
	i, j := 0, len(nums)-1
	for i < j {
		mid, count := i+(j-i)/2, 0
		for _, n := range nums {
			if n <= mid {
				count++
			}
		}
		if count > mid {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return i
}

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"7,9,7,4,2,8,7,7,1,5", args{[]int{7, 9, 7, 4, 2, 8, 7, 7, 1, 5}}, 7},
		{"1,3,4,2,2", args{[]int{1, 3, 4, 2, 2}}, 2},
		{"3,1,3,4,2", args{[]int{3, 1, 3, 4, 2}}, 3},
		{"1,1", args{[]int{1, 1}}, 1},
		{"1,1,2", args{[]int{1, 1, 2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
