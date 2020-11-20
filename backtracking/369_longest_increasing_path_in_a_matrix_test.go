package backtracking

import (
	"math"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 329. Longest Increasing Path in a Matrix
// simple backtracking, TLE
func longestIncreasingPathA(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	res := 1
	for i := range matrix {
		for j := range matrix[i] {
			lipHelperA(matrix, i, j, 1, ' ', &res)
		}
	}
	return res
}

func lipHelperA(matrix [][]int, i, j, tmp int, dir rune, res *int) {
	cur := matrix[i][j]
	up, down, left, right := math.MinInt32, math.MinInt32, math.MinInt32, math.MinInt32
	hasUp := i > 0
	if hasUp {
		up = matrix[i-1][j]
	}
	hasLeft := j > 0
	if hasLeft {
		left = matrix[i][j-1]
	}
	hasDown := i < len(matrix)-1
	if hasDown {
		down = matrix[i+1][j]
	}
	hasRight := j < len(matrix[0])-1
	if hasRight {
		right = matrix[i][j+1]
	}

	if dir == 'u' && up <= cur && left <= cur && right <= cur {
		if tmp > *res {
			*res = tmp
		}
		return
	}
	if dir == 'd' && down <= cur && left <= cur && right <= cur {
		if tmp > *res {
			*res = tmp
		}
		return
	}
	if dir == 'l' && up <= cur && down <= cur && left <= cur {
		if tmp > *res {
			*res = tmp
		}
		return
	}
	if dir == 'r' && up <= cur && down <= cur && right <= cur {
		if tmp > *res {
			*res = tmp
		}
		return
	}
	if hasUp && up > cur {
		lipHelperA(matrix, i-1, j, tmp+1, 'u', res)
	}
	if hasLeft && left > cur {
		lipHelperA(matrix, i, j-1, tmp+1, 'l', res)
	}
	if hasDown && down > cur {
		lipHelperA(matrix, i+1, j, tmp+1, 'd', res)
	}
	if hasRight && right > cur {
		lipHelperA(matrix, i, j+1, tmp+1, 'r', res)
	}
}

// backtracking with memory
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	res := 1
	dirs := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	m, n := len(matrix), len(matrix[0])
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
	}
	for i := range matrix {
		for j := range matrix[0] {
			tmp := lipHelper(matrix, i, j, m, n, cache, dirs)
			res = common.MaxInt(tmp, res)
		}
	}
	return res
}

func lipHelper(matrix [][]int, i, j, m, n int, cache [][]int, dirs [][2]int) int {
	if cache[i][j] != 0 {
		return cache[i][j]
	}
	max := 1
	for _, dir := range dirs {
		x, y := i+dir[0], j+dir[1]
		if x < 0 || x >= m || y < 0 || y >= n || matrix[x][y] <= matrix[i][j] {
			continue
		}
		tmp := 1 + lipHelper(matrix, x, y, m, n, cache, dirs)
		max = common.MaxInt(tmp, max)
	}
	cache[i][j] = max
	return max
}

func Test_lip(t *testing.T) {
	a := [][]int{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, {19, 18, 17, 16, 15, 14, 13, 12, 11, 10}, {20, 21, 22, 23, 24, 25, 26, 27, 28, 29}, {39, 38, 37, 36, 35, 34, 33, 32, 31, 30}, {40, 41, 42, 43, 44, 45, 46, 47, 48, 49}, {59, 58, 57, 56, 55, 54, 53, 52, 51, 50}, {60, 61, 62, 63, 64, 65, 66, 67, 68, 69}, {79, 78, 77, 76, 75, 74, 73, 72, 71, 70}, {80, 81, 82, 83, 84, 85, 86, 87, 88, 89}, {99, 98, 97, 96, 95, 94, 93, 92, 91, 90}, {100, 101, 102, 103, 104, 105, 106, 107, 108, 109}, {119, 118, 117, 116, 115, 114, 113, 112, 111, 110}, {120, 121, 122, 123, 124, 125, 126, 127, 128, 129}, {139, 138, 137, 136, 135, 134, 133, 132, 131, 130}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
	t.Log(longestIncreasingPath(a))
}
