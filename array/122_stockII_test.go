package array

import (
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 122. Best Time to Buy and Sell Stock II
func maxProfitII1(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp
func maxProfitII2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][2]int, len(prices))

	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][0]
}

func Test_maxProfitII(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{[]int{7, 1, 5, 3, 6, 4}}, 7},
		{"0", args{[]int{1, 2, 3, 4, 5}}, 4},
		{"0", args{[]int{5, 4, 3, 2, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitII1(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

// review 20201007
// 当天不持有：max(前一天没有当天也不买，前一天持有当天卖出)
// 当天持有：max(前一天没有当天买入，前一天持有当天不买)
func maxProfitII(prices []int) int {
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = common.MaxInt(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = common.MaxInt(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][0]
}

func Test_maxProfitII1(t *testing.T) {
	a := []int{1, 2}
	maxProfitII(a)
}
