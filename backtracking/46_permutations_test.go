package backtracking

import (
	"testing"
)

// 46. Permutations

func permute(nums []int) [][]int {
	var res [][]int
	var cur []int
	size := len(nums)
	dfsPermute(nums, cur, &res, size)
	return res
}

func dfsPermute(nums, cur []int, res *[][]int, size int) {
	for i, v := range nums {
		cur = append(cur, v)
		if len(cur) == size {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			*res = append(*res, tmp)
			return
		}
		// head := make([]int, len(nums[:i]))
		// tail := make([]int, len(nums[i+1:]))
		// copy(head, nums[:i])
		// copy(tail, nums[i+1:])
		// dfsPermute(append(head, tail...), cur, res, size)
		// ! 一次分配内存，速度比两次更快
		tmp := make([]int, len(nums)-1)
		copy(tmp[:i], nums[:i])
		copy(tmp[i:], nums[i+1:])
		dfsPermute(tmp, cur, res, size)
		cur = cur[:len(cur)-1]

	}
}

// ! 虽然有遍历，但是遍历的slice很小，然而分配内存次数少，所以性能好
func permute1(nums []int) [][]int {
	var res [][]int
	var cur []int
	dfsPermute1(nums, cur, &res)
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

func dfsPermute1(nums, cur []int, res *[][]int) {
	if len(cur) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for _, e := range nums {
		if !containsIntList(cur, e) {
			cur = append(cur, e)
			dfsPermute1(nums, cur, res)
			cur = cur[:len(cur)-1]
		}
	}
}

func Test_permute(t *testing.T) {
	a := []int{1}
	b := []int{0, 1}
	c := []int{1, 2, 3}
	t.Log(permute(a))
	t.Log(permute(b))
	t.Log(permute(c))

}

func Benchmark_permute(b *testing.B) {
	a := []int{0, 1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		permute(a)
	}
}
