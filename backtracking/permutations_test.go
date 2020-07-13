package backtracking

import (
	"fmt"
	"testing"
)

// 46. Permutations

func permute(nums []int) [][]int {
	var res [][]int
	var cur []int
	dfsPermute(nums, cur, &res)
	return res
}

func containsIntList(l []int, e int) bool {
	for _, x := range l {
		if x == e {
			return true
		}
	}
	return false
}

func dfsPermute(nums, cur []int, res *[][]int) {
	if len(cur) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for _, e := range nums {
		if !containsIntList(cur, e) {
			cur = append(cur, e)
			dfsPermute(nums, cur, res)
			cur = cur[:len(cur)-1]
		}
	}
}

func TestPermute(t *testing.T) {
	a := []int{1, 2, 3}
	b := permute(a)
	fmt.Println(b)
}
