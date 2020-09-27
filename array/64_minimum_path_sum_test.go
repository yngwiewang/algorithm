package array

import (
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 64. Minimum Path Sum

// dp 机器人寻路
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += common.MinInt(grid[i-1][j], grid[i][j-1])
		}
	}
	// fmt.Println(grid)
	return grid[m-1][n-1]
}

// 都放在一次循环里会更快一点
func minPathSum1(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && j > 0 {
				grid[i][j] += common.MinInt(grid[i-1][j], grid[i][j-1])
			} else if i == 0 && j > 0 {
				grid[0][j] += grid[0][j-1]
			} else if j == 0 && i > 0 {
				grid[i][0] += grid[i-1][0]
			}
		}
	}
	return grid[m-1][n-1]
}

func Test_minPathSum(t *testing.T) {
	a := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	// a := [][]int{}
	b := minPathSum1(a)
	t.Log(b)
}
