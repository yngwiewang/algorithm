package array

import (
	"sort"
	"testing"
)

// 90. Subsets II
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	findSubsetsII(nums, []int{}, &res, 0)
	return res
}

func findSubsetsII(n, tmp []int, res *[][]int, last int) {
	t := make([]int, len(tmp))
	copy(t, tmp)
	*res = append(*res, t)

	for i := 0; i < len(n); i++ {
		if i > 0 && n[i] == n[i-1] {
			continue
		}
		findSubsetsII(n[i+1:], append(tmp, n[i]), res, last+1)
	}
}

func Test_subsetsWithDup(t *testing.T) {
	a := []int{1, 2, 2, 2}
	res := subsetsWithDup(a)
	t.Log(res)
}
