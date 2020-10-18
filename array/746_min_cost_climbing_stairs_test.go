package array

import (
	"fmt"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 746. Min Cost Climbing Stairs
// dp
func minCostClimbingStairsA(cost []int) int {
	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = common.MinInt(dp[i-1]+cost[i], dp[i-2]+cost[i])
	}
	fmt.Println(dp)
	return common.MinInt(dp[len(cost)-2], dp[len(cost)-1])
}

// dp inplace
func minCostClimbingStairs(cost []int) int {

	for i := 2; i < len(cost); i++ {
		cost[i] = common.MinInt(cost[i-1], cost[i-2]) + cost[i]
	}
	fmt.Println(cost)
	return common.MinInt(cost[len(cost)-2], cost[len(cost)-1])
}

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"10,15", args{[]int{10, 15}}, 10},
		{"10,15,20", args{[]int{10, 15, 20}}, 15},
		{"1, 100, 1, 1, 1, 100, 1, 1, 100, 1", args{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
