package array

import (
	"reflect"
	"sort"
	"testing"
)

// 39. Combination Sum

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)
	findCombinationSum(candidates, target, 0, []int{}, &res)
	return res
}

func findCombinationSum(candidates []int, target, index int, tmp []int, res *[][]int) {
	for i := 0; i < len(candidates); i++ {
		if target == candidates[i] {
			hit := make([]int, len(tmp)+1)        // ! 只声明不行，必须初始化 hit 才能 copy
			copy(hit, append(tmp, candidates[i])) // ! copy 是因为 tmp 的变化会带进 res 里
			*res = append(*res, hit)
			return
		}
		if target > candidates[i] {
			findCombinationSum(candidates[i:], target-candidates[i], i, append(tmp, candidates[i]), res)
		}
	}
}

// * 优化 index， 递归时 candidates 不变，减小内存消耗
// ! 这种想法错了，因为 candidates 是 slice， candidates[i:] 仍然传的是 slice 里的指针
func findCombinationSum1(candidates []int, target, index int, tmp []int, res *[][]int) {
	for i := index; i < len(candidates); i++ {
		if target == candidates[i] {
			hit := make([]int, len(tmp)+1)        // ! 只声明不行，必须初始化 hit 才能 copy
			copy(hit, append(tmp, candidates[i])) // ! copy 是因为 tmp 的变化会带进 res 里
			*res = append(*res, hit)
			return
		}
		if target > candidates[i] {
			findCombinationSum1(candidates, target-candidates[i], i, append(tmp, candidates[i]), res)
		}
	}
}

func Benchmark_combinationSum(b *testing.B) {
	a := []int{1, 4, 3, 2, 23, 8, 7, 9, 5, 6}
	target := 7
	for i := 0; i < b.N; i++ {
		combinationSum(a, target)
	}
}

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1,2,4,3", args{[]int{1, 2, 4, 3}, 7}, [][]int{{2, 2, 3}, {7}}},
		{"2,3,6,7", args{[]int{2, 3, 6, 7}, 7}, [][]int{{2, 2, 3}, {7}}},
		{"2,3,5", args{[]int{2, 3, 5}, 8}, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{"2,3", args{[]int{2, 3}, 2}, [][]int{{2}}},
		{"2,3", args{[]int{2, 3}, 1}, [][]int{{}}},
		{"2", args{[]int{2}, 2}, [][]int{{2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
