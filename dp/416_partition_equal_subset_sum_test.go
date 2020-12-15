package dp

import (
	"fmt"
	"testing"
)

// 416. Partition Equal Subset Sum

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	size := len(nums)
	dp := make([][]bool, size+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}
	for i := 1; i <= size; i++ {
		for j := 1; j <= sum; j++ {
			if j-nums[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}
	return dp[size][sum]
}

func Test_canPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1,5,11,5", args{[]int{1, 5, 11, 5}}, true},
		{"1,5,2,3", args{[]int{1, 5, 2, 3}}, false},
		{"1,2,5", args{[]int{1, 2, 5}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.args.nums); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
