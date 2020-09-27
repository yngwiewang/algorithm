package array

import (
	"fmt"
	"testing"
)

// 63. Unique Paths II

// 重置地图，然后dp。注意一些边界情况要考虑
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 1
	}
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				obstacleGrid[i][j] = 1
				continue
			}
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = -1
			} else if i == 0 && j > 0 && obstacleGrid[i][j-1] == 1 {
				obstacleGrid[i][j] = 1
			} else if j == 0 && i > 0 && obstacleGrid[i-1][j] == 1 {
				obstacleGrid[i][j] = 1
			}
		}
	}
	for _, v := range obstacleGrid {
		fmt.Println(v)
	}
	fmt.Println()

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == -1 {
				continue
			}
			if obstacleGrid[i-1][j] != -1 {
				obstacleGrid[i][j] += obstacleGrid[i-1][j]
			}
			if obstacleGrid[i][j-1] != -1 {
				obstacleGrid[i][j] += obstacleGrid[i][j-1]
			}
		}
	}
	for _, v := range obstacleGrid {
		fmt.Println(v)
	}

	return obstacleGrid[m-1][n-1]
}

// 参考答案简化代码
func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	obstacleGrid[0][0] = 1
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 && obstacleGrid[i-1][0] == 1 {
			obstacleGrid[i][0] = 1
		} else {
			obstacleGrid[i][0] = 0
		}
	}
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 0 && obstacleGrid[0][i-1] == 1 {
			obstacleGrid[0][i] = 1
		} else {
			obstacleGrid[0][i] = 0
		}
	}
	for _, v := range obstacleGrid {
		fmt.Println(v)
	}
	fmt.Println()

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}
	for _, v := range obstacleGrid {
		fmt.Println(v)
	}

	return obstacleGrid[m-1][n-1]
}

func Test_uniquePathsWithObstacles(t *testing.T) {
	// a := [][]int{{0, 1, 0, 0, 0}, {1, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
	a := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	b := uniquePathsWithObstacles1(a)
	fmt.Println(b)
}
