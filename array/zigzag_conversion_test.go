package array

import (
	"fmt"
	"strings"
	"testing"
)

// 6. ZigZag Conversion
// * 注意 mod 0 的情况
// * 如果字符串的元素不止是 ascii 字母，要注意转换成rune
func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}
	period := 2*numRows - 2
	res := make([]string, numRows)

	for i, v := range s {
		mod := i % period
		if mod < numRows {
			res[mod] += string(v)
		} else {
			res[period-mod] += string(v)
		}
	}
	return strings.Join(res, "")
}

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"Abc", 1}, "Abc"},
		{"", args{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{"", args{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test(t *testing.T) {
	a := 'a'
	b := byte(a)
	fmt.Printf("%T,%s  %T,%c\n", b, string(b), a, a)

}
