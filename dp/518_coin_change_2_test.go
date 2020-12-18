package dp

import (
	"fmt"
	"testing"
)

// 518. Coin Change 2
// dp table
func change2(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}
	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if coins[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}
	return dp[len(coins)][amount]
}

// 压缩存储
func change2a(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if coins[i] <= j {
				dp[j] = dp[j] + dp[j-coins[i]]
			}
			fmt.Println(dp)
		}
	}
	return dp[amount]
}

func Test_change2(t *testing.T) {
	type args struct {
		amount int
		coins  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1,2,5 5", args{5, []int{1, 2, 5}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := change2a(tt.args.amount, tt.args.coins); got != tt.want {
				t.Errorf("change2() = %v, want %v", got, tt.want)
			}
		})
	}
}
