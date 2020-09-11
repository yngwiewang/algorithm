package string

import (
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

//856. Score of Parentheses

// 使用栈记录括号包含的层次，时空复杂度都是O(n)
func scoreOfParentheses(S string) int {
	level := make([]int, 1)
	level[0] = 0

	for _, s := range S {
		if s == '(' {
			level = append(level, 0)
		} else {
			a := level[len(level)-1]
			b := level[len(level)-2]
			level = level[:len(level)-2]
			level = append(level, b+common.MaxInt(2*a, 1))
		}
	}
	return level[len(level)-1]
}

// 记录括号对的个数，判断嵌套的成以2，时间O(n)，空间O(1)
func scoreOfParentheses1(S string) int {
	res, count := 0, 0
	for i, s := range S {
		if s == '(' {
			count++
		} else {
			count--
			if S[i-1] == '(' {
				res += 1 << count
			}
		}
	}
	return res
}

func Test_scoreOfParentheses(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"()", args{"()"}, 1},
		{"()()", args{"()()"}, 2},
		{"(()(()))", args{"(()(()))"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreOfParentheses1(tt.args.S); got != tt.want {
				t.Errorf("scoreOfParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
