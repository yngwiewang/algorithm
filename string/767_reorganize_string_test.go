package string

import (
	"fmt"
	"testing"
)

// 767. Reorganize String
// 参考最好的答案做的，很巧妙，不用排序，时空复杂度都战胜100%
func reorganizeString(S string) string {
	var cc [26]int
	for _, v := range S {
		cc[v-'a']++
	}
	pos, max := maxChar(&cc)
	if max > (len(S)+1)/2 {
		return ""
	}
	res := make([]byte, len(S))
	idx := 0
	for cc[pos] > 0 {
		res[idx] = byte(pos) + 'a'
		idx += 2
		cc[pos]--
	}
	for i := 0; i < 26; i++ {
		for cc[i]>0{
			if idx >= len(res){
				idx = 1
			}
			res[idx] = byte(i) + 'a'
			idx += 2
			cc[i]--
		}
	}
	return string(res)
}

func maxChar(cc *[26]int) (pos, max int) {
	for i, v := range cc {
		if v > max {
			pos, max = i, v
		}
	}
	return
}

func Test_reorganizeString(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"aabbccdd", args{"aabbccdd"}, "acacbdbd"},
		{"aab", args{"aab"}, "aba"},
		{"aaab", args{"aaab"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorganizeString(tt.args.S); got != tt.want {
				t.Errorf("reorganizeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
