package dp

import (
	"fmt"
	"testing"
)

// 647. Palindromic Substrings

func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				res++
			}
			for _, v := range dp {
				fmt.Println(v)
			}
			fmt.Println()
		}
	}
	return res
}

func countSubstrings1(s string) int {
	res := 0
	for i := range s {
		res += countSubstringHelper(s, i, i) + countSubstringHelper(s, i, i+1)
	}
	return res
}

func countSubstringHelper(s string, i, j int) int {
	count := 0
	for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {
		count++
	}
	return count
}

func Test_countSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"abc", args{"abc"}, 3},
		{"aaa", args{"aaa"}, 6},
		{"fdsklf", args{"fdsklf"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubstrings1(tt.args.s); got != tt.want {
				t.Errorf("countSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
