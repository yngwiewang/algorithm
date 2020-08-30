package string

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

// 3. Longest Substring Without Repeating Characters

func lengthOfLongestSubstring(s string) int {
	longest := 0
	substring := ""
	for i := 0; i < len(s); i++ {
		if strings.Contains(substring, string(s[i])) {
			substring = substring[strings.Index(substring, string(s[i]))+1:]
		}
		substring += string(s[i])
		if len(substring) > longest {
			longest = len(substring)
		}
	}
	return longest
}

func lengthOfLongestSubstring1(s string) int {
	longest := 0
	substring := ""
	for i := 0; i < len(s); i++ {
		if strings.Contains(substring, string(s[i])) {
			substring = substring[strings.Index(substring, string(s[i]))+1:]
		}
		substring += string(s[i])
		if len(substring) > longest {
			longest = len(substring)
		}
	}
	return longest
}

// 涉及中文要用 utf8.RuneLen 和  utf8.RuneCountInString
func lengthOfLongestSubstring2(s string) int {
	longest := 0
	substring := ""
	for _, c := range s {
		if strings.Contains(substring, string(c)) {
			substring = substring[strings.Index(substring, string(c))+ utf8.RuneLen(c):]
		}
		substring += string(c)
		fmt.Println(substring)
		if utf8.RuneCountInString(substring) > longest {
			longest = len(substring)
		}
	}
	return longest
}

// ! leetcode fastest
func lengthOfLongestSubstring3(s string) int {
    best := 0    
    start := 0
    for i, a := range s {
        for j, b := range s[start:i] {
            if a == b && i - start > best {
                best = i - start
                start = start + j + 1
                break
            } else if a == b {
                start = start + j + 1
                break
            }
        }
    }
    
    if len(s) - start > best {
        best = len(s) - start
    }
    
    return best
}

func Test_lengthOfLongestSubstring1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"abcdabcbb", args{"abcabcbb"}, 3},
		{"abcabcbb", args{"abcabcbb"}, 3},
		{"aaabb", args{"aaabb"}, 2},
		{"bbbbbbbbb", args{"bbbbbbbbb"}, 1},
		{"", args{""}, 0},
		{" ", args{" "}, 1},
		{"中国人中国", args{"中国人中国"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring3(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
