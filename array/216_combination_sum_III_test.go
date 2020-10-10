package array

import "testing"

// 216. Combination Sum III
// backtracing
func combinationSumIII(k int, n int) [][]int {
	res := make([][]int, 0)
	findCombinationSumIII(k, n, 1, 0, &res, []int{})
	return res
}

func findCombinationSumIII(k, n, idx, sum int, res *[][]int, tmp []int) {
	if sum == n && len(tmp) == k {
		t := make([]int, len(tmp))
		copy(t, tmp)
		*res = append(*res, t)
		return
	}
	if sum > n || len(tmp) > k {
		return
	}
	for i := idx; i <= 9; i++ {
		findCombinationSumIII(k, n, i+1, sum+i, res, append(tmp, i))
	}
}

func Test_findCombinationSumIII(t *testing.T) {
	for _, v := range [][2]int{{3, 7}, {3, 9}, {4, 1}, {3, 2}, {9, 45}, {2, 1}, {2, 60}} {
		t.Log(v)
		k, n := v[0], v[1]
		res := combinationSumIII(k, n)
		t.Log(res)
	}

}
