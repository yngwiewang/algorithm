package backtracking

import (
	"fmt"
	"testing"
)

// 518. Coin Change 2

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, c := range coins {
		for j := c; j < amount+1; j++ {
			dp[j] += dp[j-c]
			fmt.Println(c, j, dp)
		}
	}
	return dp[amount]
}

func Test_change(t *testing.T) {
	type args struct {
		amount int
		coins  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"5 1,2,5", args{5, []int{1, 2, 5}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := change(tt.args.amount, tt.args.coins); got != tt.want {
				t.Errorf("change() = %v, want %v", got, tt.want)
			}
		})
	}
}
