package slidingwindow

import (
	"reflect"
	"testing"
)

// 239. Sliding Window Maximum

// 窗口保存索引
func maxSlidingWindow1(nums []int, k int) []int {
	// res := make([]int, (len(nums)+1)-k)
	// win := make([]int, 0, k)
	var res []int
	var win []int
	for i, v := range nums {
		if i >= k && win[0] == i-k {
			win = win[1:]
		}
		for len(win) > 0 && nums[win[len(win)-1]] <= v {
			win = win[:len(win)-1]
		}
		win = append(win, i)
		if i+1 >= k {
			// res[i-k+1] = nums[win[0]]
			res = append(res, nums[win[0]])
		}
	}
	return res
}

// 窗口保存值，单调队列
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, (len(nums)+1)-k)
	win := make([]int, 0, k)
	for i, v := range nums {
		for len(win) > 0 && win[len(win)-1] < v {
			win = win[:len(win)-1]
		}
		win = append(win, v)
		if i >= k && win[0] == nums[i-k] {
			win = win[1:]
		}
		if i+1 >= k {
			res[i-k+1] = win[0]
			// res = append(res, win[0])
		}
	}
	return res
}

func Benchmark_maxSlidingWindow(b *testing.B) {
	a := []int{9, 10, 9, -7, -4, -8, 2, -6}

	for i := 0; i < b.N; i++ {
		_ = maxSlidingWindow1(a, 5)
	}
}

// 暴力
func maxSlidingWindow2(nums []int, k int) []int {
	res := make([]int, (len(nums)+1)-k)
	max := -10000

	for i := 0; i <= len(nums)-k; i++ {
		if max > -10000 && nums[i-1] != max && nums[i+k-1] < max {
			res[i] = max
			continue
		} else {
			max = -10000
		}
		for _, v := range nums[i : i+k] {
			if v > max {
				max = v
			}
		}
		res[i] = max
	}
	return res
}

func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"2", args{[]int{1, 3, 1, 2, 0, 5}, 3}, []int{3, 3, 2, 5}},
		{"1", args{[]int{7, 2, 4}, 2}, []int{7, 4}},
		{"3", args{[]int{9, 10, 9, -7, -4, -8, 2, -6}, 5}, []int{10, 10, 9, 2}},
		{"4", args{[]int{-7, -8, 7, 5, 7, 1, 6, 0}, 4}, []int{7, 7, 7, 7, 7}},
		{"5", args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3}, []int{3, 3, 5, 5, 6, 7}},
		{"6", args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 2}, []int{3, 3, -1, 5, 5, 6, 7}},
		{"7", args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 4}, []int{3, 5, 5, 6, 7}},
		{"8", args{[]int{1}, 1}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
