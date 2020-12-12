package dp

import (
	"fmt"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 516. Longest Palindromic Subsequence

func longestPalindromeSubseq(s string) int {
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(s))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	res := lpsHelper(s, 0, len(s)-1, memo)
	return res
}

func lpsHelper(s string, i, j int, memo [][]int) int {
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if i == j {
		memo[i][j] = 1
		return 1
	}
	if i > j {
		memo[i][j] = 0
		return 0
	}
	if s[i] == s[j] {
		memo[i][j] = 2 + lpsHelper(s, i+1, j-1, memo)
	} else {
		memo[i][j] = common.MaxInt(lpsHelper(s, i+1, j, memo), lpsHelper(s, i, j-1, memo))
	}
	return memo[i][j]
}

func longestPalindromeSubseq1(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		// for j := range dp[i] {
		// 	dp[i][j] = 1
		// }
	}
	for i := len(s) - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = 2 + dp[i+1][j-1]
			} else {
				dp[i][j] = common.MaxInt(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}
	return dp[0][len(s)-1]
}

func longestPalindromeSubseq20201211(s string) int {
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(s))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	res := lpssHelper20201211(s, 0, len(s)-1, memo)
	fmt.Println(memo)
	return res
}

func lpssHelper20201211(s string, i, j int, memo [][]int) int {
	if i > j {
		return 0
	}
	if i == j {
		return 1
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if s[i] == s[j] {
		memo[i][j] = 2 + lpssHelper20201211(s, i+1, j-1, memo)
		return memo[i][j]
	}
	memo[i][j] = common.MaxInt(lpssHelper20201211(s, i+1, j, memo),
		lpssHelper20201211(s, i, j-1, memo))
	return memo[i][j]
}

func longestPalindromeSubseq20201211dp(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}

	for i := len(s) - 2; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = 2 + dp[i+1][j-1]
			} else {
				dp[i][j] = common.MaxInt(dp[i][j-1], dp[i+1][j])
			}
		}
	}

	return dp[0][len(s)-1]
}

func Test_longestPalindromeSubseq(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"bbbab", args{"bbbab"}, 4},
		{"cbbd", args{"cbbd"}, 2},
		{"a", args{"a"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeSubseq20201211dp(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
