package array

import (
	"math"
	"sort"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 16. 3Sum Closest

// 双指针，照着solution写的，O(n2)，效率低
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	min := 10e4
	for i := 0; i < len(nums)-2; i++ {
		low, high := i+1, len(nums)-1
		for low < high {
			sum := nums[i] + nums[low] + nums[high]
			if math.Abs(float64(target-sum)) < math.Abs(min) {
				min = float64(target - sum)
			}
			if sum < target {
				low++
			} else {
				high--
			}
		}
	}
	return target - int(min)
}

// 双指针，按照最好答案优化，主要就是自己写一个求整数绝对距离的函数
func threeSumClosest1(nums []int, target int) int {
	sort.Ints(nums)
	min := 10000
	res := 0
	for i := 0; i < len(nums)-2; i++ {
		low, high := i+1, len(nums)-1
		for low < high {
			sum := nums[i] + nums[low] + nums[high]
			if sum < target {
				low++
			} else {
				high--
			}
			if common.DistanceInt(target, sum) < min {
				min = common.DistanceInt(target, sum)
				res = sum
			}
		}
	}
	return res
}

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1,1,-1,-1,3", args{[]int{1, 1, -1, -1, 3}, -1}, -1},
		{"-1,2,1,-4,3,-9,0,7,-2,1,7,5,2", args{[]int{-1, 2, 1, -4, 3, -9, 0, 7, -2, 1, 7, 5, 2}, 10}, 10},
		{"-1,2,1,-4", args{[]int{-1, 2, 1, -4}, 1}, 2},
		{"0,2,1,-3", args{[]int{0, 2, 1, -3}, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest1(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
