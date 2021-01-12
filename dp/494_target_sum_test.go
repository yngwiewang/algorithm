package dp

import (
	"fmt"
	"testing"
)

// 494. Target Sum

func findTargetSumWays1(nums []int, S int) int {
	// memo := map[[2]int]int{}
	res := 0
	findTargetsSumWaysHelper1(nums, &res, 0, 0, S)
	return res

}

func findTargetsSumWaysHelper1(nums []int, res *int, i, sum, s int) {
	if i == len(nums) {
		if sum == s {
			*res++
		}
		return
	}
	findTargetsSumWaysHelper1(nums, res, i+1, sum+nums[i], s)
	findTargetsSumWaysHelper1(nums, res, i+1, sum-nums[i], s)
}

// recursive with memo
func findTargetSumWays2(nums []int, S int) int {
	memo := map[[2]int]int{}
	return findTargetsSumWaysHelper2(nums, 0, S, memo)
}

func findTargetsSumWaysHelper2(nums []int, i, s int, memo map[[2]int]int) int {
	if i == len(nums) {
		if s == 0 {
			return 1
		}
		return 0
	}
	if v, ok := memo[[2]int{i, s}]; ok {
		return v
	}
	res := findTargetsSumWaysHelper2(nums, i+1, s+nums[i], memo) +
		findTargetsSumWaysHelper2(nums, i+1, s-nums[i], memo)
	memo[[2]int{i, s}] = res
	return res
}

// dp, how many subsets make sum(subsets) == (target + sum(nums)) / 2
func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if (S+sum)%2 != 0 {
		return 0
	}
	s := (S + sum) / 2
	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, s+1)
		dp[i][0] = 1
	}
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= s; j++ {
			dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}
	return dp[len(nums)][s]
}

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums []int
		S    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 1, 1, 1, 1}, 3}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.S); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
