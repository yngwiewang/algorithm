package array

import (
	"fmt"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 714. Best Time to Buy and Sell Stock with Transaction Fee
// dp
func maxProfitWithFeeA(prices []int, fee int) int {
	dp := make([][2]int, len(prices))
	dp[0][0], dp[0][1] = 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = common.MaxInt(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = common.MaxInt(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	for _, v := range dp {
		fmt.Println(v)
	}
	return dp[len(prices)-1][0]
}

// dp with less space but more time
func maxProfitWithFeeB(prices []int, fee int) int {
	dp := [2][2]int{}
	dp[0][0], dp[0][1] = 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i%2][0] = common.MaxInt(dp[(i-1)%2][0], dp[(i-1)%2][1]+prices[i]-fee)
		dp[i%2][1] = common.MaxInt(dp[(i-1)%2][0]-prices[i], dp[(i-1)%2][1])
	}
	for _, v := range dp {
		fmt.Println(v)
	}
	return dp[(len(prices)-1)%2][0]
}

// dp with less space and less time
// ! best
func maxProfitWithFee(prices []int, fee int) int {
	sold, hold := 0, -prices[0]

	for i := 1; i < len(prices); i++ {
		// tmp := hold
		sold = common.MaxInt(sold, hold+prices[i]-fee)
		hold = common.MaxInt(sold-prices[i], hold)
		fmt.Println(hold, sold)
	}

	return sold
}

func Test_maxProfitWithFee(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1, 3, 2, 8, 4, 9", args{[]int{1, 3, 2, 8, 4, 9}, 2}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitWithFee(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfitWithFee() = %v, want %v", got, tt.want)
			}
		})
	}
}
