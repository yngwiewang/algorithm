package dp

import (
	"math"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 322. Coin Change
// dp, recursive with memo
func coinChange(coins []int, amount int) int {
	var memo = make(map[int]int)

	res := ccHelper(coins, amount, memo)
	return res
}

func ccHelper(coins []int, amount int, memo map[int]int) int {
	if tmp, ok := memo[amount]; ok {
		return tmp
	}
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res := math.MaxInt32
	for _, coin := range coins {
		tmp := ccHelper(coins, amount-coin, memo)
		if tmp == -1 {
			continue
		}
		res = common.MinInt(res, 1+tmp)
	}
	if res == math.MaxInt32 {
		memo[amount] = -1
	} else {
		memo[amount] = res
	}
	return memo[amount]
}

// dp table
func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for i := range dp {
		for _, coin := range coins {
			if i >= coin { // ! 这里应用了cpu指令控制器的分支预测，显著提高性能。
				dp[i] = common.MinInt(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[len(dp)-1] == amount+1 {
		return -1
	}
	return dp[len(dp)-1]
}

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"6", args{[]int{1, 2, 5}, 100}, 20},
		{"5", args{[]int{1}, 2}, 2},
		{"2", args{[]int{2}, 3}, -1},
		{"0", args{[]int{1, 2, 5}, 11}, 3},
		{"1", args{[]int{1, 2, 5}, 12}, 3},
		{"3", args{[]int{1}, 0}, 0},
		{"4", args{[]int{1}, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange1(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
