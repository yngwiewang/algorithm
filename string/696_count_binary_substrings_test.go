package string

import (
	"math"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 696. Count Binary Substrings

// 逐个判断，性能很差
func countBinarySubstrings1(s string) int {
	res := 0
	for i := range s {
		cntZero, cntOne := 0, 0
		j := i
		if s[j] == '0' {
			for j < len(s) && s[j] == '0' {
				cntZero, j = cntZero+1, j+1
			}
			if j == len(s) {
				continue
			}
			for j < len(s) && s[j] == '1' {
				cntOne, j = cntOne+1, j+1
				if cntOne == cntZero {
					res++
					break
				}
			}
		} else if s[j] == '1' {
			for j < len(s) && s[j] == '1' {
				cntOne, j = cntOne+1, j+1
			}
			if j == len(s) {
				continue
			}
			for j < len(s) && s[j] == '0' {
				cntZero, j = cntZero+1, j+1
				if cntOne == cntZero {
					res++
					break
				}
			}
		}
	}
	return res
}

// 0、1分组，从第二组开始加较小值
func countBinarySubstrings2(s string) int {
	l := []int{1}
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			l[len(l)-1]++
		} else {
			l = append(l, 1)
		}
	}
	res := 0
	for i := 1; i < len(l); i++ {
		res += common.MinInt(l[i], l[i-1])
	}
	return res
}

// 0、1分组，从第二组开始累j较小值，不使用额外数组
// 时空复杂度达到最小
func countBinarySubstrings3(s string) int {
	prev, cur, res := 0, 1, 0

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cur++
		} else {
			res += common.MinInt(prev, cur)
			prev = cur
			cur = 1
		}
	}
	return res + common.MinInt(prev, cur)
}

// 有少量答案使用内置的比较函数达到100%，测试无效
func countBinarySubstrings(s string) int {
	l := []float64{1.0}
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			l[len(l)-1]++
		} else {
			l = append(l, float64(1))
		}
	}
	res := 0.0
	for i := 1; i < len(l); i++ {
		res += math.Min(l[i], l[i-1])
	}
	return int(res)
}

func Test_countBinarySubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"10", args{"10"}, 1},
		{"01", args{"01"}, 1},
		{"1", args{"1"}, 0},
		{"0", args{"0"}, 0},
		{"10101", args{"10101"}, 4},
		{"00110011", args{"00110011"}, 6},
		{"01", args{"01"}, 1},
		{"001100", args{"001100"}, 4},
		{"00110", args{"00110"}, 3},
		{"0011", args{"0011"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBinarySubstrings(tt.args.s); got != tt.want {
				t.Errorf("countBinarySubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
