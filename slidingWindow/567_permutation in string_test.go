package slidingwindow

import (
	"testing"
)

// 567. Permutation in String

func checkInclusion1(s1 string, s2 string) bool {
	need, window := make(map[byte]int, 26), make(map[byte]int, 26)
	for _, v := range s1 {
		need[byte(v)]++
	}
	l, r := 0, 0
	valid := 0
	for r < len(s2) {
		if _, ok := need[s2[r]]; ok {
			window[s2[r]]++
			if window[s2[r]] == need[s2[r]] {
				valid++
			}
		}
		r++
		for r-l == len(s1) {
			if valid == len(need) {
				return true
			}
			if _, ok := need[s2[l]]; ok {
				if window[s2[l]] == need[s2[l]] {
					valid--
				}
				window[s2[l]]--
			}
			l++
		}
	}
	return false
}

// improve, more faster
func checkInclusion(s1 string, s2 string) bool {
	m1, m2 := [26]uint8{}, [26]uint8{}
	for _, v := range s1 {
		m1[v-'a']++
	}
	for l, r := 0, 0; r < len(s2); r++ {
		m2[s2[r]-'a']++
		if r-l+1 == len(s1) {
			if m1 == m2 {
				return true
			}
			m2[s2[l]-'a']--
			l++
		}
	}
	return false
}

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"4", args{"b", "ab"}, true},
		{"5", args{"ab", "eidboaoo"}, false},
		{"9", args{"trinitrophenylmethylnitramine", "dinitrophenylhydrazinetrinitrophenylmethylnitramine"}, true},
		{"8", args{"abcdxabcde", "aabcdeabcdx"}, true},
		{"7", args{"abcdxabcde", "abcdeabcdx"}, true},
		{"6", args{"abc", "bbbca"}, true},
		{"3", args{"a", "ab"}, true},
		{"2", args{"a", "a"}, true},
		{"1", args{"ab", "eidbaooo"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
