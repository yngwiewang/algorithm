package array

import (
	"math"
	"testing"
)

// 209. Minimum Size Subarray Sum
// tow pointers, sliding window.
func minSubArrayLen1(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j := 0, 0
	sum := nums[i]
	minLen := math.MaxInt32
	for {
		if sum < s {
			j++
			if j == len(nums) {
				break
			}
			sum += nums[j]
		} else {
			if j-i+1 < minLen {
				minLen = j - i + 1
			}
			sum -= nums[i]
			i++
		}
	}
	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

// answer, shorter
func minSubArrayLen(s int, nums []int) int {
	var sum, i, j, minLen int
	for j = 0; j < len(nums); j++ {
		sum += nums[j]
		for i <= j && sum >= s {
			tmp := j - i + 1
			if minLen == 0 {
				minLen = tmp
			} else if minLen > tmp {
				minLen = tmp
			}
			sum -= nums[i]
			i++
		}
	}
	return minLen
}

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		s    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{2, []int{}}, 0},
		{"2", args{2, []int{2}}, 1},
		{"2", args{3, []int{2}}, 0},
		{"2,3,1,2,4,3", args{50, []int{2, 3, 1, 2, 4, 3}}, 0},
		{"2,3,1,2,4,3", args{7, []int{2, 3, 1, 2, 4, 3}}, 2},
		{"1,1,1,2", args{2, []int{1, 1, 1, 2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.s, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
