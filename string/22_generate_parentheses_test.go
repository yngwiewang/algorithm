package string

import (
	"reflect"
	"testing"
)

// 22. Generate Parentheses

func generateParenthesis(n int) []string {
	res := &[]string{}
	if n == 0 {
		return *res
	}
	addParenthes(res, "", 0, 0, n)
	return *res
}

func addParenthes(res *[]string, tmp string, left, right, n int) {
	if len(tmp) == n*2 {
		*res = append(*res, tmp)
		return
	}
	if left < n {
		addParenthes(res, tmp+"(", left+1, right, n)
	}
	if right < left {
		addParenthes(res, tmp+")", left, right+1, n)
	}
}

// ! 使用字节切片可以节省内存消耗，快了3倍
func generateParenthesis1(n int) []string {
	res := &[]string{}
	if n == 0 {
		return *res
	}
	bytes := make([]byte, 0, 2*n)
	addParenthes1(res, bytes, 0, 0, n)
	return *res
}

func addParenthes1(res *[]string, tmp []byte, left, right, n int) {
	if len(tmp) == n*2 {
		*res = append(*res, string(tmp))
		return
	}
	if left < n {
		addParenthes1(res, append(tmp, '('), left+1, right, n)
	}
	if right < left {
		addParenthes1(res, append(tmp, ')'), left, right+1, n)
	}
}

func Benchmark_gp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := 5
		_ = generateParenthesis1(n)
	}
}

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"", args{3}, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{"", args{0}, []string{}},
		{"", args{1}, []string{"()"}},
		{"", args{2}, []string{"(())", "()()"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis1(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
