package backtracking

import (
	"testing"
)

// 77. Combinations

func combine(n int, k int) [][]int {
	res := [][]int{}
	cur := make([]int, 0, k)
	combineHelper(n, 1, k, cur, &res)
	return res
}

func combineHelper(n, p, k int, cur []int, res *[][]int) {
	if len(cur) == k {
		tmp := make([]int, k)
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := p; i <= n; i++ {
		combineHelper(n, i+1, k, append(cur, i), res)
	}
}

func Test_combine(t *testing.T) {
	n, k := 5, 3
	t.Log(combine(n, k))
}
