package string

import (
	"math"
	"reflect"
	"strconv"
	"testing"
)

// 842. Split Array into Fibonacci Sequence

// 递归、贪心、回溯
// 时间复杂度100%
// ! 小心题设要求32位整数
func splitIntoFibonacci(S string) []int {
	res := []int{}
	if len(S) < 3 {
		return res
	}

	res = btArrayToFib(S, []int{})
	return res
}

func btArrayToFib(s string, fib []int) (res []int) {
	if len(res) > 0 {
		return
	}
	if len(fib) > 2 && fib[len(fib)-1] != fib[len(fib)-2]+fib[len(fib)-3] {
		return
	}
	if len(s) == 0 {
		if len(fib) > 2 && fib[len(fib)-1] == fib[len(fib)-2]+fib[len(fib)-3] {
			res = fib
		}
		return
	}
	if s[0] == '0' {
		res = btArrayToFib(s[1:], append(fib, 0))
	} else {
		for i := 1; i <= len(s) && len(res) == 0; i++ {
			num, _ := strconv.Atoi(s[:i])
			if num > 2147483647 {
				return
			}
			res = btArrayToFib(s[i:], append(fib, num))
		}
	}
	return
}

func Test_splitIntoFibonacci(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"74912134825162255812723932620170946950766784234934", args{"74912134825162255812723932620170946950766784234934"}, []int{}},
		{"1101111", args{"1101111"}, []int{11, 0, 11, 11}},
		{"000", args{"000"}, []int{0, 0, 0}},
		{"0123", args{"0123"}, []int{}},
		{"112358130", args{"112358130"}, []int{}},
		{"11235813", args{"11235813"}, []int{1, 1, 2, 3, 5, 8, 13}},
		{"123", args{"123"}, []int{1, 2, 3}},
		{"123456579", args{"123456579"}, []int{123, 456, 579}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitIntoFibonacci(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitIntoFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max32(t *testing.T) {
	a := math.MaxInt32
	t.Log(a)

	b := "2147483647"
	c, ok := strconv.Atoi(b)
	t.Log(c, ok)
}
