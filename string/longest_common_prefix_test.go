package string

import (
	"strings"
	"testing"
)

// 14. Longest Common Prefix

// 优化方法是不要累加common，而是直接赋值
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	common := ""
	i := 0
	for ; i < len(strs[0]); i++ {
		count := 0
		for _, v := range strs {
			if len(v) == i {
				return common
			}
			if v[i] == strs[0][i] {
				count++
			}
		}
		if count == len(strs) {
			common = strs[0][:i+1]
		} else {
			return common
		}

	}
	return common
}

func minStr(strs []string) string {
	min := strs[0]
	for _, v := range strs {
		if v < min {
			min = v
		}
	}
	return min
}

func maxStr(strs []string) string {
	max := strs[0]
	for _, v := range strs {
		if v > max {
			max = v
		}
	}
	return max
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	min, max := minStr(strs), maxStr(strs)
	for i, v := range min {
		if byte(v) != max[i] {
			return min[:i]
		}
	}
	return min
}

// 小浩算法给出的方法，也比较快，不额外消耗内存。
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	common := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) == 0 {
			return ""
		}
		for strings.Index(strs[i], common) != 0 {

			common = common[:len(common)-1]
		}

	}
	return common
}

// 这个方法是最快的，也没有额外的内存分配
func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	prefix := strs[0]
	i := 0
	for ; i < len(prefix); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != prefix[i] {
				return prefix[:i]
			}
		}
	}
	return prefix[:i]
}

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0.", args{[]string{"aa", "a"}}, "a"},
		{"1.", args{[]string{"aa", "aa"}}, "aa"},
		{"2.", args{[]string{}}, ""},
		{"3.", args{[]string{""}}, ""},
		{"4.", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"5.", args{[]string{"dog", "dog", "cat"}}, ""},
		{"6.", args{[]string{"aa", "aa", "aa"}}, "aa"},
		{"7.", args{[]string{"aa", "", "aa"}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix3(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0.", args{[]string{"aa", "a"}}, "a"},
		{"0.", args{[]string{"aa", "aa"}}, "aa"},
		{"0.", args{[]string{}}, ""},
		{"0.", args{[]string{""}}, ""},
		{"1.", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"2.", args{[]string{"dog", "dog", "cat"}}, ""},
		{"3.", args{[]string{"aa", "aa", "aa"}}, "aa"},
		{"4.", args{[]string{"aa", "", "aa"}}, ""},
	}
	for i := 0; i < b.N; i++ {
		for _, t := range tests {
			longestCommonPrefix1(t.args.strs)
		}
	}
}
