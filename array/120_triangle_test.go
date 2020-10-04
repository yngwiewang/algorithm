package array

import (
	"math"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 20. Triangle
// 回溯 TLE
func minimumTotal1(triangle [][]int) int {
	res := math.MaxInt32
	findMinimumTotal(triangle, 0, 0, 0, &res)
	return res
}

func findMinimumTotal(triangle [][]int, tmp, level, idx int, res *int) {
	if level == len(triangle)-1 {
		if tmp+triangle[level][idx] < *res {
			*res = tmp + triangle[level][idx]
		}
		return
	}
	findMinimumTotal(triangle, tmp+triangle[level][idx], level+1, idx, res)
	findMinimumTotal(triangle, tmp+triangle[level][idx], level+1, idx+1, res)
}

// 自顶向下逐层循环
func minimumTotal2(triangle [][]int) int {
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			a, b := math.MaxInt32, math.MaxInt32
			if j > 0 {
				a = triangle[i-1][j-1]
			}
			if j < len(triangle[i-1]) {
				b = triangle[i-1][j]
			}
			triangle[i][j] = triangle[i][j] + common.MinInt(a, b)
		}
	}
	return common.MinIntInSlice(triangle[len(triangle)-1])
}

// 自底向上逐层循环
func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = triangle[i][j] + common.MinInt(
				triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func Test_minimumTotal(t *testing.T) {
	a := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	res := minimumTotal(a)
	t.Log(res)
}
