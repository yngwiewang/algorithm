package array

import (
	"testing"
)

// 867. Transpose Matrix
func transpose(a [][]int) [][]int {
	r, c := len(a), len(a[0])
	res := make([][]int, c)
	for i := 0; i < c; i++ {
		res[i] = make([]int, r)
	}
	for i, v := range res {
		for j := range v {
			res[i][j] = a[j][i]
		}
	}
	return res
}

func Test_transpose(t *testing.T) {
	// a := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	// a := [][]int{{1, 2}}
	a := [][]int{{1}}
	b := transpose(a)
	t.Log(b)
}
