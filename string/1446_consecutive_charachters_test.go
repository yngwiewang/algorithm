package string

import "testing"

// 1446. Consecutive Characters

func maxPower(s string) int {
	power := 0
	for i := 0; i < len(s); {
		tmp, j := 0, i
		for ; j < len(s) && s[i] == s[j]; j++ {
			tmp++
		}
		if tmp > power {
			power = tmp
		}
		i = j
	}
	return power
}

func maxPower1(s string) int {
	c, mp := 0, 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			c = c + 1
		} else {
			if c > mp {
				mp = c
			}
			c = 0
		}
	}
	if c > mp {
		mp = c
	}
	return mp + 1
}

func Benchmark_maxPower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maxPower("abbcccddddeeeeedcba")
	}
}

func Test_maxPower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"leetcode", args{"leetcode"}, 2},
		{"abbcccddddeeeeedcba", args{"abbcccddddeeeeedcba"}, 5},
		{"triplepillooooow", args{"triplepillooooow"}, 5},
		{"hooraaaaaaaaaaay", args{"hooraaaaaaaaaaay"}, 11},
		{"tourist", args{"tourist"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPower1(tt.args.s); got != tt.want {
				t.Errorf("maxPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
