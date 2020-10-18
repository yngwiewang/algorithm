package array

import "testing"

// 766. Toeplitz Matrix
// 逐行比较，节省空间
func isToeplitzMatrixA(matrix [][]int) bool {
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}

// leetcode 上更快一点，benchmark 并没有得到验证
func isToeplitzMatrix(matrix [][]int) bool {
	for i, m := range matrix {
		for j, n := range m {
			if i == 0 || j == 0 {
				continue
			}
			if n != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}

func Test_isToeplitzMatrix(t *testing.T) {
	// a := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	a := [][]int{{22, 0, 94, 45, 46, 96}, {10, 22, 80, 94, 45, 46}}
	t.Log(isToeplitzMatrix(a))
}

func Benchmark_isTeoplitzMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
		isToeplitzMatrix(a)
	}
}
