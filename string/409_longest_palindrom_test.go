package string

import (
	"fmt"
	"testing"
)

// 409. Longest Palindrome

func longestPalindrome(s string) int {
	count := make(map[rune]int)
	for _, v := range s {
		count[v]++
	}
	fmt.Println(count)
	res := 0
	hasEven := false
	for _, v := range count {
		res += v / 2 * 2

		if v%2 == 1 {
			hasEven = true
		}
	}
	if hasEven {
		res++
	}
	return res
}

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"abccccdd", args{"abccccdd"}, 7},
		{"a", args{"a"}, 1},
		{"bb", args{"bb"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
