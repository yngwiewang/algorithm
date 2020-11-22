package db

import (
	"fmt"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 1143. Longest Common Subsequence
// dp, recursive
func longestCommonSubsequence1(text1 string, text2 string) int {
	memo := make(map[[2]int]int)
	res := dpLCS1(text1, 0, text2, 0, memo)
	return res
}

func dpLCS1(s1 string, i int, s2 string, j int, memo map[[2]int]int) int {
	if i == len(s1) || j == len(s2) {
		return 0
	}
	if tmp, ok := memo[[2]int{i, j}]; ok {
		return tmp
	}
	if s1[i] == s2[j] {
		memo[[2]int{i, j}] = 1 + dpLCS1(s1, i+1, s2, j+1, memo)
		return memo[[2]int{i, j}]
	}
	memo[[2]int{i, j}] = common.MaxInt(dpLCS1(s1, i+1, s2, j, memo),
		dpLCS1(s1, i, s2, j+1, memo))
	return memo[[2]int{i, j}]
}

// use []byte instead of string and [][]int instead of map[[2]int]int,
// the latter makes it more faster.
func longestCommonSubsequence2(text1 string, text2 string) int {
	memo := make([][]int, len(text1))
	for i := 0; i < len(text1); i++ {
		memo[i] = make([]int, len(text2))
		for j := 0; j < len(text2); j++ {
			memo[i][j] = -1
		}
	}
	res := dpLCS2([]byte(text1), 0, []byte(text2), 0, memo)
	return res
}

func dpLCS2(b1 []byte, i int, b2 []byte, j int, memo [][]int) int {
	if i == len(b1) || j == len(b2) {
		return 0
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if b1[i] == b2[j] {
		memo[i][j] = 1 + dpLCS2(b1, i+1, b2, j+1, memo)
		return memo[i][j]
	}
	memo[i][j] = common.MaxInt(dpLCS2(b1, i+1, b2, j, memo),
		dpLCS2(b1, i, b2, j+1, memo))
	return memo[i][j]
}

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(text1)+1; i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = common.MaxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}
	return dp[len(text1)][len(text2)]
}

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"abcde, ace", args{"abcde", "ace"}, 3},
		{"abcde, ae", args{"abcde", "ae"}, 2},
		{"abc, abc", args{"abc", "abc"}, 3},
		{"abc, def", args{"abc", "def"}, 0},
		{"a, a", args{"a", "a"}, 1},
		{"a, b", args{"a", "b"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
