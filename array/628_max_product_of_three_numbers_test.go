package array

import (
	"math"
	"sort"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 628. Maximum Product of Three Numbers
// sort O(nLogn)
func maximumProductA(nums []int) int {
	sort.Ints(nums)
	return common.MaxInt(nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3],
		nums[0]*nums[1]*nums[len(nums)-1])
}

// find min1,min2,max1,max2,max3
// O(n)
func maximumProduct(nums []int) int {
	min1, min2, max1, max2, max3 := math.MaxInt32, math.MaxInt32, math.MinInt32, math.MinInt32, math.MinInt32
	for _, v := range nums {
		if v < min2 {
			if v < min1 {
				min2, min1 = min1, v
			} else {
				min2 = v
			}
		}
		if v > max1 {
			if v > max2 {
				if v > max3 {
					max1, max2, max3 = max2, max3, v
				} else {
					max1, max2 = max2, v
				}
			} else {
				max1 = v
			}
		}
	}
	return common.MaxInt(min1*min2*max3, max1*max2*max3)
}

func Test_maximumProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"-4,0,5,6", args{[]int{-4, 0, 5, 6}}, 0},
		{"1,2,3,4", args{[]int{1, 2, 3, 4}}, 24},
		{"-1,0,1,2", args{[]int{-1, 0, 1, 2}}, 0},
		{"-1,-2,0,1", args{[]int{-1, -2, 0, 1}}, 2},
		{"-1,-2,-3,0", args{[]int{-1, -2, -3, 0}}, 0},
		{"-1,-2,-3,0,2", args{[]int{-1, -2, -3, 0, 2}}, 12},
		{"-1,-2,-3,-4", args{[]int{-1, -2, -3, -4}}, -6},
		{"-1,-2,-3,-4,5", args{[]int{-1, -2, -3, -4, 5}}, 60},
		{"-1,-2,-3,4,5", args{[]int{-1, -2, -3, 4, 5}}, 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProduct(tt.args.nums); got != tt.want {
				t.Errorf("maximumProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
