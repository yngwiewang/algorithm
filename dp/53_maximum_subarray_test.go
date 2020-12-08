package dp

import (
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 53. Maximum Subarray
// dp
func maxSubArray1(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = common.MaxInt(nums[i], nums[i]+dp[i-1])
		res = common.MaxInt(res, dp[i])
	}
	return res
}

func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = common.MaxInt(nums[i], nums[i]+nums[i-1])
		res = common.MaxInt(res, nums[i])
	}
	return res
}

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1}}, 1},
		{"0", args{[]int{0}}, 0},
		{"-1", args{[]int{-1}}, -1},
		{"-2147483647", args{[]int{-2147483647}}, -2147483647},
		{"-2,1,-3,4,-1,2,1,-5,4", args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
