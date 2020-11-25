package dp

import (
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 712. Minimum ASCII Delete Sum for Two Strings

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = dp[0][i-1] + int(s2[i-1])
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = common.MinInt(int(s1[i-1])+dp[i-1][j], int(s2[j-1])+dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
func Test_minimumDeleteSum(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sea, eat", args{"sea", "eat"}, 231},
		{"delete, leet", args{"delete", "leet"}, 403},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeleteSum(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("minimumDeleteSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
