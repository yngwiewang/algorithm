package array

import "testing"

// 59. Spiral Matrix II

// 与 54 题类似，从 [0,-1] 初始，按四个方向遍历
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n)
	}
	col, row := n, n-1
	a, b := 0, -1
	direction := right
	for i := 1; i <= n*n; {
		switch direction {
		case right:
			for j := 0; j < col; j++ {
				b++
				res[a][b] = i
				i++
			}
			col, direction = col-1, down
		case down:
			for j := 0; j < row; j++ {
				a++
				res[a][b] = i
				i++
			}
			row, direction = row-1, left
		case left:
			for j := 0; j < col; j++ {
				b--
				res[a][b] = i
				i++
			}
			col, direction = col-1, up
		case up:
			for j := 0; j < row; j++ {
				a--
				res[a][b] = i
				i++
			}
			row, direction = row-1, right
		}
	}
	return res
}

func Test_generateMatrix(t *testing.T) {
	a := 4
	b := generateMatrix(a)
	for _, v := range b {
		t.Log(v)
	}
}
