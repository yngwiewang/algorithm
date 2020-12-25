package binarysearch

import "testing"

// 33. Search in Rotated Sorted Array

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}
		if nums[l] <= nums[m] {
			if nums[m] > target && nums[l] <= target {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if nums[m] < target && nums[r] >= target {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1
}

func binarySearch(n []int, t int) int {
	if len(n) == 0 {
		return -1
	}
	l, r := 0, len(n)-1
	for l < r-1 {
		m := l + (r-l)/2
		if n[m] == t {
			return m
		} else if n[m] < t {
			l = m + 1
		} else {
			r = m
		}
	}
	if n[l] == t {
		return l
	}
	return -1
}

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"8", args{[]int{5, 4}, 4}, 1},
		{"12", args{[]int{5, 1, 3}, 3}, 2},
		{"5", args{[]int{3, 4, 5, 0, 1, 2}, 2}, 5},
		{"3", args{[]int{3, 4, 5, 0, 1, 2}, 0}, 3},
		{"11", args{[]int{1, 3}, 3}, 1},
		{"10", args{[]int{1, 3}, 0}, -1},
		{"0", args{[]int{3, 4, 5, 0, 1, 2}, 3}, 0},
		{"1", args{[]int{3, 4, 5, 0, 1, 2}, 4}, 1},
		{"2", args{[]int{3, 4, 5, 0, 1, 2}, 5}, 2},
		{"4", args{[]int{3, 4, 5, 0, 1, 2}, 1}, 4},
		{"6", args{[]int{3, 4, 5, 0, 1, 2}, 6}, -1},
		{"7", args{[]int{5, 4}, 5}, 0},
		{"9", args{[]int{1}, 0}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
