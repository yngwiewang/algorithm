package dp

import (
	"fmt"
	"testing"
)

// 5. Longest Palindromic Substring
func longestPalindrome(s string) string {
	if s == "" || len(s) == 1 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	l, r := 0, 0
	for i := len(s) - 2; i >= 0; i-- {
		dp[i][i] = true
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
			}
			if dp[i][j] && j-i > r-l {
				l, r = i, j
			}
			fmt.Println(dp)
			fmt.Println(i, j, l, r)
		}
	}
	return s[l : r+1]
}

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"aa", args{"aa"}, "a"},
		{"a", args{"a"}, "a"},
		{"babad", args{"babad"}, "aba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
