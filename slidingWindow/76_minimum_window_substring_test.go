package slidingwindow

import (
	"math"
	"testing"
)

// 76. Minimum Window Substring

func minWindow1(s string, t string) string {
	need, window := make(map[byte]int, len(t)), make(map[byte]int, len(t))
	minLen := math.MaxInt32
	res := ""
	for _, v := range t {
		need[byte(v)]++
		window[byte(v)] = 0
	}
	l, r := 0, 0

	for ; l < len(s); l++ {
		for ; r < len(s); r++ {
			coverd := true
			if _, ok := need[s[r]]; ok {
				window[s[r]]++
				for k := range need {
					if window[k] < need[k] {
						coverd = false
						break
					}
				}
			} else {
				continue
			}
			if !coverd {
				if r == len(s)-1 {
					return res
				}
				continue
			}

			for ; l < len(s); l++ {
				coverd = true
				if _, ok := need[s[l]]; ok {
					window[s[l]]--
					for k := range need {
						if window[k] < need[k] {
							coverd = false
							break
						}
					}
				} else {
					continue
				}
				if coverd == true {
					continue
				}
				if r-l < minLen {
					minLen = r - l
					res = s[l : r+1]
				}
				l++
				break
			}
		}
	}
	return res
}

// improved a little
func minWindow(s string, t string) string {
	need, window := make(map[byte]int, len(t)), make(map[byte]int, len(t))
	minLen := math.MaxInt32
	res := ""
	for _, v := range t {
		need[byte(v)]++
		window[byte(v)] = 0
	}
	l, r := 0, 0

	for ; l < len(s); l++ {
		for ; r < len(s); r++ {
			if _, ok := need[s[r]]; !ok {
				continue
			}
			coverd := true
			window[s[r]]++
			for k := range need {
				if window[k] < need[k] {
					coverd = false
					break
				}
			}
			if !coverd {
				if r == len(s)-1 {
					return res
				}
				continue
			}
		OUT:
			for ; l < len(s); l++ {
				if _, ok := need[s[l]]; ok {
					window[s[l]]--
					if window[s[l]] < need[s[l]] {
						if r-l < minLen {
							minLen = r - l
							res = s[l : r+1]
						}
						l++
						break OUT
					}
				}
			}
		}
	}
	return res
}

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"acbdbaab", "aabd"}, "dbaa"},
		{"1", args{"ADOBECODEBANC", "ABC"}, "BANC"},
		{"1", args{"a", "a"}, "a"},
		{"1", args{"a", "k"}, ""},
		{"1", args{"abdsfsdfcab", "adk"}, ""},
		{"1", args{"abdsfsdfcab", "ac"}, "ca"},
		{"1", args{"abdsfsdfcab", "ac"}, "ca"},
		{"1", args{"abdsfsdfca", "ac"}, "ca"},
		{"1", args{"abca", "ac"}, "ca"},
		{"1", args{"abc", "ac"}, "abc"},
		{"1", args{"ad", "ad"}, "ad"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
