package dp

import (
	"fmt"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 198. House Robber
// dp table
func rob(nums []int) int {
	dp := make([]int, len(nums)+2)
	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = common.MaxInt(dp[i+1], nums[i]+dp[i+2])
	}
	fmt.Println(dp)
	return dp[0]
}

func rob1(nums []int) int {
	dp_i_1, dp_i_2 := 0, 0
	dp_i := 0
	for i := len(nums) - 1; i >= 0; i-- {
		dp_i = common.MaxInt(dp_i_1, nums[i]+dp_i_2)
		dp_i_2, dp_i_1 = dp_i_1, dp_i
	}
	return dp_i
}

// dp recursive
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}
	robHelper(nums, 0, memo)
	return memo[0]
}

func robHelper(nums []int, i int, memo []int) int {
	if i >= len(nums) {
		return 0
	}
	if memo[i] != -1 {
		return memo[i]
	}
	memo[i] = common.MaxInt(nums[i]+robHelper(nums, i+2, memo),
		robHelper(nums, i+1, memo))
	return memo[i]
}

func rob20201218(nums []int) int {
	dpi, dp1, dp2 := 0, 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		dpi = common.MaxInt(nums[i]+dp2, dp1)
		dp2, dp1 = dp1, dpi
	}
	return dp1
}

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1,2,3,1", args{[]int{1, 2, 3, 1}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob20201218(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
