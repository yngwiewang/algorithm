package string

import (
	"reflect"
	"testing"
)

// 17. Letter Combinations of a Phone Number
// * 递归回溯
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	dict := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	res := []string{}
	level := ""
	dealLevel(&res, &level, digits, dict)
	return res
}

func dealLevel(res *[]string, level *string, digits string, dict map[string][]string) {
	if len(digits) == 0 {
		*res = append(*res, *level)
		return
	}
	for i := 0; i < len(dict[string(digits[0])]); i++ {
		*level += dict[string(digits[0])][i]
		dealLevel(res, level, digits[1:], dict)
		*level = (*level)[:len(*(level))-1]
	}
}

// * leetcode的方法，更快，减少内存消耗
func letterCombinations1(digits string) []string {
	if digits == "" {
		return []string{}
	}
	dict := map[int][]string{
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}

	res := []string{}
	level := ""
	dealLevel1(&res, &level, digits, dict)
	return res
}

func dealLevel1(res *[]string, level *string, digits string, dict map[int][]string) {
	if len(digits) == 0 {
		*res = append(*res, *level)
		return
	}
	for i := 0; i < len(dict[int(digits[0]-'0')]); i++ {
		*level += dict[int(digits[0]-'0')][i]
		dealLevel1(res, level, digits[1:], dict)
		*level = (*level)[:len(*(level))-1]
	}
}

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"2,3", args{"546"}, []string{"jgm","jgn","jgo","jhm","jhn","jho","jim","jin","jio","kgm","kgn","kgo","khm","khn","kho","kim","kin","kio","lgm","lgn","lgo","lhm","lhn","lho","lim","lin","lio"}},
		{"2,3", args{"23"}, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinations1(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCombinations1("23")
	}
}
