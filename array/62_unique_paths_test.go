package array

import (
	"fmt"
	"testing"
)

// 62. Unique Paths

// 动态规划，把首行和首列置1，
// 然后走到每个格子的步数等于走到左侧格子的步数+走到上侧格子的步数
func uniquePaths(m int, n int) int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		res[0][i] = 1
	}
	for i := 1; i < m; i++ {
		res[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[i][j] = res[i-1][j] + res[i][j-1]
		}
	}
	for _, v := range res {
		fmt.Println(v)
	}
	return res[m-1][n-1]
}

func Test_uniquePaths(t *testing.T) {
	m, n := 3, 3
	res := uniquePaths(m, n)
	t.Log(res)
}
