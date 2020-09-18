package array

import (
	"reflect"
	"sort"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 40. Combination Sum II
// 回溯，比39题简单的地方在于原数组中的数字不能重复使用，
// 复杂的地方在于也在于此，因为结果中不能包含重复的切片
// 使用简单的判断重复切片的方法，不到最快结果
func combinationSum2(candidates []int, target int) [][]int {
	var res = [][]int{}
	sort.Ints(candidates)
	findCombinationSum2(candidates, target, 0, []int{}, &res)
	return res
}

func findCombinationSum2(candidates []int, target, index int, tmp []int, res *[][]int) {
	if target == 0 {
		hit := make([]int, len(tmp))
		copy(hit, tmp)
		for _, v := range *res {
			if common.EqualInts(v, hit) {
				return
			}
		}
		*res = append(*res, hit)
		return
	}
	for i := index; i < len(candidates); i++ {
		if target >= candidates[i] {
			findCombinationSum2(candidates, target-candidates[i], i+1, append(tmp, candidates[i]), res)
		} else {
			break
		}
	}
}

// 参考最快答案，剪枝优化，100%
func combinationSum3(candidates []int, target int) [][]int {
	var res = [][]int{}
	sort.Ints(candidates)
	findCombinationSum3(candidates, target, 0, []int{}, &res)
	return res
}

func findCombinationSum3(candidates []int, target, index int, tmp []int, res *[][]int) {
	if target == 0 {
		hit := make([]int, len(tmp))
		copy(hit, tmp)
		*res = append(*res, hit)
		return
	}
	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue // ! 剪枝优化
		}
		if target >= candidates[i] { // ! 注意这里要有等号
			findCombinationSum3(candidates, target-candidates[i], i+1, append(tmp, candidates[i]), res)
		} else {
			break // ! 剪枝优化
		}
	}
}

func Benchmark_combinationSum2(b *testing.B) {
	a := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	for i := 0; i < b.N; i++ {
		combinationSum2(a, target)
	}
}

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"10,1,2,7,6,1,5", args{[]int{10, 1, 2, 7, 6, 1, 5}, 8}, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{"2,5,2,1,2", args{[]int{2, 5, 2, 1, 2}, 5}, [][]int{{1, 2, 2}, {5}}},
		{"", args{[]int{}, 5}, [][]int{}},
		{"1", args{[]int{1}, 5}, [][]int{}},
		{"1", args{[]int{1}, 1}, [][]int{{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum3(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
