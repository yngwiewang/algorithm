package dp

import (
	"fmt"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 300. Longest Increasing Subsequence
// dp tablel, O(n2)
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	res := 1
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = common.MaxInt(dp[i], dp[j]+1)
				res = common.MaxInt(dp[i], res)
			}
		}
	}
	fmt.Println(dp)
	return res
}

func lengthOfLIS20201206(nums []int) int {
	dp := make([]int, len(nums))
	res := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = common.MaxInt(dp[i], dp[j]+1)
			}
			// ! 如果把res的比较放在这里要注意res的初始值是1而不能是0
			res = common.MaxInt(res, dp[i]) 
		}
	}
	fmt.Println(dp)
	return res
}

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"[0]", args{[]int{0}}, 1},
		{"[1,3,6,7,9,4,10,5,6]", args{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6}}, 6},
		{"[10,9,2,5,3,7,101,18]", args{[]int{10, 9, 2, 5, 3, 7, 101, 18}}, 4},
		{"[0,1,0,3,2,3]", args{[]int{0, 1, 0, 3, 2, 3}}, 4},
		{"[7,7,7,7,7,7,7]", args{[]int{7, 7, 7, 7, 7, 7, 7}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS20201206(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
