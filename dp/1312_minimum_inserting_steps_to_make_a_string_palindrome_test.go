package dp

import (
	"fmt"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 1312. Minimum Insertion Steps to Make a String Palindrome
// dp, recursive
func minInsertions1(s string) int {
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(s))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	res := minInsertionsHelper(s, 0, len(s)-1, memo)
	for _, v := range memo {
		fmt.Println(v)
	}
	return res

}

func minInsertionsHelper(s string, i, j int, memo [][]int) int {
	if i == j {
		return 0
	}
	if i > j {
		return -1
	}
	if i+1 == j && s[i] == s[j] {
		memo[i][j] = 0
		return memo[i][j]
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if s[i] == s[j] {
		memo[i][j] = minInsertionsHelper(s, i+1, j-1, memo)
		return memo[i][j]
	}
	memo[i][j] = 1 + common.MinInt(minInsertionsHelper(s, i+1, j, memo),
		minInsertionsHelper(s, i, j-1, memo))
	return memo[i][j]
}

// dp table
func minInsertions(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
	}
	for i := len(s) - 2; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = 1 + common.MinInt(dp[i+1][j],
					dp[i][j-1])
			}
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}
	return dp[0][len(s)-1]
}

func Test_minInsertions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"zzazz", args{"zzazz"}, 0},
		{"mbadm", args{"mbadm"}, 2},
		{"g", args{"g"}, 0},
		{"no", args{"no"}, 1},
		{"leetcode", args{"leetcode"}, 5},
		{"zjveiiwvc", args{"zjveiiwvc"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minInsertions(tt.args.s); got != tt.want {
				t.Errorf("minInsertions() = %v, want %v", got, tt.want)
			}
		})
	}
}
