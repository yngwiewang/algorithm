package slidingwindow

import (
	"reflect"
	"testing"
)

// 438. Find All Anagrams in a String

func findAnagrams1(s string, p string) []int {
	need, window := make(map[byte]int), make(map[byte]int)
	res := []int{}
	for _, v := range p {
		need[byte(v)]++
	}
	l := 0
	matched := 0
	for r, c := range s {
		if _, ok := need[byte(c)]; !ok {
			continue
		}
		window[byte(c)]++
		if window[byte(c)] == need[byte(c)] {
			matched++
		}
		for r-l+1 >= len(p) {
			if matched == len(need) && r-l+1 == len(p) {
				res = append(res, l)
			}
			if _, ok := need[byte(s[l])]; ok {
				if window[s[l]] == need[s[l]] {
					matched--
				}
				window[s[l]]--
			}
			l++
		}
	}
	return res
}

// use array, more faster
func findAnagrams(s string, p string) []int {
	window, need := [128]byte{}, [128]byte{}
	for _, v := range p {
		need[byte(v)]++
	}
	res := []int{}
	l := 0
	for r, c := range s {
		window[byte(c)]++
		for r-l+1 >= len(p) {
			if window == need {
				res = append(res, l)
			}
			window[byte(s[l])]--
			l++
		}
	}
	return res
}

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"5", args{"aaabaaaa", "aaa"}, []int{0, 4, 5}},
		{"4", args{"acdcaeccde", "c"}, []int{1, 3, 6, 7}},
		{"3", args{"aababaa", "aa"}, []int{0, 5}},
		{"2", args{"cbaebabacd", "abc"}, []int{0, 6}},
		{"1", args{"abab", "ab"}, []int{0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
