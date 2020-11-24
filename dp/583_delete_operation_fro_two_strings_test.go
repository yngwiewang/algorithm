package dp

import (
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 583. Delete Operation for Two Strings
// len(word1) + len(word2) - 2 * len(longest common string)
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = common.MaxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	lcs := dp[m][n]
	return m + n - 2*lcs
}

func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sea, eat", args{"sea", "eat"}, 2},
		{"abcde, ace", args{"abcde", "ace"}, 3},
		{", ", args{"", ""}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
