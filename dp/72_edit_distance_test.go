package dp

import (
	"fmt"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 72. Edit Distance
// dp, recursive
func minEditDistance(word1 string, word2 string) int {
	memo := make([][]int, len(word1))
	for i := range memo {
		memo[i] = make([]int, len(word2))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	res := minEditDistanceHelper(word1, word2, len(word1)-1, len(word2)-1, memo)
	for _, v := range memo {
		fmt.Println(v)
	}
	return res
}

func minEditDistanceHelper(w1, w2 string, i, j int, memo [][]int) int {
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if w1[i] == w2[j] {
		memo[i][j] = minEditDistanceHelper(w1, w2, i-1, j-1, memo)
		return memo[i][j]
	}
	memo[i][j] = 1 + common.MinIntInSlice([]int{minEditDistanceHelper(w1, w2, i-1, j, memo),
		minEditDistanceHelper(w1, w2, i, j-1, memo),
		minEditDistanceHelper(w1, w2, i-1, j-1, memo)})

	return memo[i][j]
}

// dp table
func minEditDistance1(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for _, v := range dp {
		fmt.Println(v)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + common.MinIntInSlice([]int{
					dp[i-1][j], dp[i-1][j-1], dp[i][j-1]})
			}
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}

	return dp[m][n]
}

func Test_minEditDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"horse ros", args{"horse", "ros"}, 3},
		{"intention execution", args{"intention", "execution"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEditDistance1(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
