package backtracking

import (
	"fmt"
	"testing"
)

// 78. Subsets

func subsets(nums []int) [][]int {
	var (
		res [][]int
		tmp []int
	)
	backtrackSubset(nums, tmp, 0, &res)
	return res
}

func backtrackSubset(nums []int, tmp []int, start int, result *[][]int) {
	cur := make([]int, len(tmp))
	copy(cur, tmp)
	*result = append(*result, cur)
	for i := start; i < len(nums); i++ {
		tmp = append(tmp, nums[i])

		backtrackSubset(nums, tmp, i+1, result)
		tmp = tmp[:len(tmp)-1]
	}
}

func subsets1(nums []int) [][]int {
	res := [][]int{nil}
	for _, num := range nums {
		for _, l := range res {
			tmp := make([]int, len(l)+1)
			tmp[0] = num
			copy(tmp[1:], l)
			res = append(res, tmp)
		}
	}
	return res
}

func TestSubsets(t *testing.T) {
	a := []int{1, 2, 3}
	b := subsets(a)
	fmt.Println(b)
}
