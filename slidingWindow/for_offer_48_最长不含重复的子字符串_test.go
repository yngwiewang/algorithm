package slidingwindow

import "testing"

// 剑指 Offer 48. 最长不含重复字符的子字符串

func lengthOfLongestSubstring1(s string) int {
	dict := make(map[byte]int)
	l := 0
	res := 0
	for r, c := range s {
		dict[byte(c)]++
		for dict[byte(c)] > 1 {
			dict[s[l]]--
			l++
		}
		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}

// use array, beats 100%
func lengthOfLongestSubstring(s string) int {
	dict := [128]byte{}
	l := 0
	res := 0
	for r, c := range s {
		dict[c]++
		for dict[c] > 1 {
			dict[s[l]]--
			l++
		}
		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{" ", args{" "}, 1},
		{"", args{""}, 0},
		{"pwwkew", args{"pwwkew"}, 3},
		{"bbbbb", args{"bbbbb"}, 1},
		{"abcabcbb", args{"abcabcbb"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
