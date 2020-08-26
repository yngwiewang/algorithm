package string

import (
	"math"
	"testing"
)

// 387. First Unique Character in a String
func firstUniqChar(s string) int {
	dict := make(map[byte]int, 26)
	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if dict[s[i]] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar1(s string) int {
	d1 := make(map[byte]int)
	d2 := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		d1[s[i]]++
		if _, ok := d2[s[i]]; !ok {
			d2[s[i]] = i
		}
	}
	first := math.MaxInt64
	for k, v := range d1 {
		if v == 1 {
			if d2[k] < first {
				first = d2[k]
			}
		}
	}
	if first == math.MaxInt64 {
		return -1
	}
	return first
}

func firstUniqChar2(s string) int {
	seq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		seq[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if seq[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar3(s string) int {
	seq := make([]int, 26)
	for _, v := range s {
		seq[v-'a']++
	}
	for i, v := range s {
		if seq[v-'a'] == 1 {
			return i
		}
	}
	return -1
}

func BenchmarkFirstUniqChar(b *testing.B) {
	x := "jweoiurefvnnoweimrjvuirehvwoleijgfoehrvuliyrenbvciwsoervjoperhnvikujhbewvcouiherfveiavjosdhjgvifedhgbuv"
	for i := 0; i < b.N; i++ {
		firstUniqChar(x)
	}
}
func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"leetcode"}, 0},
		{"", args{"loveleetcode"}, 2},
		{"", args{"aa"}, -1},
		{"", args{""}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar3(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
