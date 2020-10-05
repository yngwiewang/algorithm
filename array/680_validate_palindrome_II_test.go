package array

import (
	"testing"
)

// 680. Valid Palindrome II
// 双指针循环，代码复杂混乱，边界判断易错，时空复杂度高
func validPalindrome1(s string) bool {
	i, j := 0, len(s)-1
	a, b := 0, 0
	skipLeft, skipRight := false, false
	for {
		if i >= j {
			return s[i] == s[j]
		}
		if s[i] != s[j] {
			if !(skipLeft || skipRight) {
				if i+1 == j || i == j-1 {
					return true
				}
				a, b = i, j
				i++
				skipLeft = true
				continue
			}
			if skipRight {
				return false
			}
			if skipLeft {
				i, j = a, b
				j--
				skipRight = true
				continue
			}
		}
		i, j = i+1, j-1
	}
}

// 找到中止回文时两个指针的位置，分别跳过每个指继续j判断是否回文
func validPalindromeSimple(s string) bool{
	for i:=0 ;i<len(s) /2;i++{
		if s[i] != s[len(s)-i-1]{
			return false
		}
	}
	return true
}

func validPalindrome(s string) bool {
	for i:=0 ;i<len(s) /2;i++{
		if s[i] != s[len(s)-i-1]{
			j := len(s)-i-1
			return validPalindromeSimple(s[i:j]) || validPalindromeSimple(s[i+1:j+1])
		}
	}
	return true
}


func Test_validPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"lcupuupucul", args{"lcupuupucul"}, true},
		{"lcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupucul", args{"lcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupucul"}, true},
		{"abc", args{"abc"}, false},
		{"abcaabca", args{"abcaabca"}, false},
		{"ab", args{"ab"}, true},
		{"abca", args{"abca"}, true},
		{"a", args{"a"}, true},
		{"aba", args{"aba"}, true},
		{"abbca", args{"abbca"}, true},
		{"ababca", args{"ababca"}, true},
		{"abaabca", args{"abaabca"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.args.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
