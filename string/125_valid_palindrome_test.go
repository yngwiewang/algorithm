package string

import (
	"strings"
	"testing"
	"unicode"
)

// 125. Valid Palindrome

// * 转换成字节数组，只保留字母和数字并转换成小写
// * 然后做回文判断
// ! 内存占用较高
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') {
			l = append(l, s[i])
		}
	}
	for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
		if l[i] != l[j] {
			return false
		}
	}
	return true
}

// * 双指针做回文判断
// ! 最快，内存最小
func isPalindrome1(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !(unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i]))) {
			i++
		} else if !(unicode.IsLetter(rune(s[j])) || unicode.IsDigit(rune(s[j]))) {
			j--
		} else {
			if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
				return false
			}
			i, j = i+1, j-1
		}
	}
	return true
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"a b C,:K"}, false},
		{"", args{"A man, a plan, a canal: Panama"}, true},
		{"", args{"0P"}, false},
		{"", args{"P"}, true},
		{"", args{"0"}, true},
		{"", args{"0,"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome1(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
