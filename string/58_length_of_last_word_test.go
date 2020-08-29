package string

import (
	"strings"
	"testing"
)

// 58. Length of Last Word

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	s = strings.Trim(s, " ")

	var count int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 32 {
			break
		}
		count++
	}
	return count
}

func lengthOfLastWord1(s string) int {
	var count int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 32 {
			if count == 0 {
				continue
			}
			break
		}
		count++
	}
	return count
}

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"  sdf ds dfsdf   sdf  dsfdsf f "}, 1},
		{"", args{"hello world"}, 5},
		{"", args{"hello world "}, 5},
		{"", args{" "}, 0},
		{"", args{"   "}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord1(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
