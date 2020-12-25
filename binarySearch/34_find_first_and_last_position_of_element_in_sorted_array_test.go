package binarysearch

import (
	"reflect"
	"testing"
)

// 34. Find First and Last Position of Element in Sorted Array

func searchRange1(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}
	find := -1
	i, j := 0, len(nums)-1
	for i <= j {
		m := i + (j-i)/2
		if nums[m] == target {
			find = m
			break
		} else if nums[m] > target {
			j = m - 1
		} else {
			i = m + 1
		}
	}
	if find == -1 {
		return res
	}
	l, r := 0, find
	for l < r {
		m := l + (r-l)/2
		if nums[m] == target {
			r = m
		} else {
			l = m + 1
		}
	}
	res[0] = l
	l, r = find, len(nums)-1
	for l < r-1 {
		m := l + (r-l)/2
		if nums[m] == target {
			l = m
		} else {
			r = m - 1
		}
	}
	if nums[r] == target {
		res[1] = r
	} else {
		res[1] = l
	}
	return res
}

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 || target < nums[0] || target > nums[len(nums)-1] {
		return res
	}
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if nums[m] == target {
			r = m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if nums[l] == target {
		res[0] = l
	} else {
		return res
	}
	l, r = 0, len(nums)-1
	for l < r-1 {
		m := l + (r-l)/2
		if nums[m] == target {
			l = m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	if nums[r] == target {
		res[1] = r
	} else {
		res[1] = l
	}
	return res
}

func TestBinarySearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	t.Log(searchRange(a, 6))
}

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{5, 7, 7, 8, 8, 10}, 8}, []int{3, 4}},
		{"", args{[]int{1, 2, 3, 8, 8, 8, 9}, 8}, []int{3, 5}},
		{"", args{[]int{1, 2, 3, 8, 8, 8, 8}, 8}, []int{3, 6}},
		{"", args{[]int{1, 2, 3, 8, 8, 8, 8}, 6}, []int{-1, -1}},
		{"", args{[]int{8}, 8}, []int{0, 0}},
		{"", args{[]int{8, 8}, 8}, []int{0, 1}},
		{"", args{[]int{8, 8, 8}, 8}, []int{0, 2}},
		{"", args{[]int{8, 8, 8, 8}, 8}, []int{0, 3}},
		{"", args{[]int{1, 8, 9}, 8}, []int{1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
