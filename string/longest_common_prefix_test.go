package string

import "testing"

// 14. Longest Common Prefix
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	common := ""
	for i := 0; i < len(strs[0]); i++ {
		count := 0
		for _, v := range strs {
			if len(v) == i{
				return common
			}
			if v[i] == strs[0][i] {
				count++
			}
		}
		if count == len(strs) {
			common += string(strs[0][i])
		} else {
			return common
		}

	}
	return common
}

func minStr(strs []string) string{
	min := strs[0]
	for _, v := range strs{
		if v < min{
			min = v
		}
	}
	return min
}

func maxStr(strs []string) string{
	max := strs[0]
	for _, v := range strs{
		if v > max{
			max = v
		}
	}
	return max
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	min,max := minStr(strs), maxStr(strs)
	for i, v := range min{
		if byte(v) != max[i]{
			return min[:i]
		}
	}
	return min
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
		{"0.", args{[]string{"aa", "aa"}}, "aa"},
		{"0.", args{[]string{}}, ""},
		{"0.", args{[]string{""}}, ""},
		{"1.", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"2.", args{[]string{"dog", "dog", "cat"}}, ""},
		{"3.", args{[]string{"aa", "aa", "aa"}}, "aa"},
		{"4.", args{[]string{"aa", "", "aa"}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
