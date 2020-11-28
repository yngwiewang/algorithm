package dp

import (
	"fmt"
	"testing"
)

// 10. Regular Expression Matching
// recursive with memo
func isMatch(s string, p string) bool {
	memo := make(map[[2]int]bool)
	res := isMatchHelper(s, p, 0, 0, memo)
	return res
}

// func isMatchHelper(s, p string, i, j int, memo map[[2]int]bool) bool {
// 	if v, ok := memo[[2]int{i, j}]; ok {
// 		return v
// 	}
// 	if j >= len(p) {
// 		memo[[2]int{i, j}] = len(s) == i
// 		return memo[[2]int{i, j}]
// 	}
// 	first := i < len(s) && (s[i] == p[j] || p[j] == '.')
// 	if j < len(p)-1 && p[j+1] == '*' {
// 		memo[[2]int{i, j}] = isMatchHelper(s, p, i, j+2, memo) || first &&
// 			isMatchHelper(s, p, i+1, j, memo)
// 		return memo[[2]int{i, j}]
// 	}
// 	memo[[2]int{i, j}] = first && isMatchHelper(s, p, i+1, j+1, memo)
// 	return memo[[2]int{i, j}]
// }
// ! 利用cpu分支预测提高性能
func isMatchHelper(s, p string, i, j int, memo map[[2]int]bool) bool {
	if v, ok := memo[[2]int{i, j}]; ok {
		return v
	}
	if j >= len(p) {
		memo[[2]int{i, j}] = len(s) == i
		return memo[[2]int{i, j}]
	}
	first := i < len(s) && (s[i] == p[j] || p[j] == '.')
	if !(j < len(p)-1 && p[j+1] == '*') {
		memo[[2]int{i, j}] = first && isMatchHelper(s, p, i+1, j+1, memo)
		return memo[[2]int{i, j}]
	}

	memo[[2]int{i, j}] = isMatchHelper(s, p, i, j+2, memo) || first &&
		isMatchHelper(s, p, i+1, j, memo)
	return memo[[2]int{i, j}]
}

// use dp table
func isMatch1(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[m][n] = true
	for i := m; i >= 0; i-- {
		for j := n-1; j >= 0; j-- {
			firstMatch := i < m && (p[j] == s[i] || p[j] == '.')
			if j+1 < n && p[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || firstMatch && dp[i+1][j]
			} else {
				dp[i][j] = firstMatch && dp[i+1][j+1]
			}
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}
	return dp[0][0]
}

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"5", args{"mississippi", "mis*is*p*."}, false},
		{"4", args{"aab", "c*a*b"}, true},
		{"0", args{"aa", "a"}, false},
		{"1", args{"aa", "a."}, true},
		{"2", args{"aa", "a*"}, true},
		{"3", args{"ab", ".*"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch1(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
