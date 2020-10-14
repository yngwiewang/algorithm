package array

import "testing"

// 566. Reshape the Matrix
// 遍历添加
func matrixReshapeA(nums [][]int, r int, c int) [][]int {
	width, height := len(nums[0]), len(nums)
	if r*c != width*height {
		return nums
	}
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	a, b := 0, 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			res[a][b] = nums[i][j]
			b++
			if b == c {
				b = 0
				a++
			}
		}
	}
	return res
}

// 添加的时候用取余和取模来定位值，巧妙
func matrixReshape(nums [][]int, r int, c int) [][]int {
	width, height := len(nums[0]), len(nums)
	if r*c != width*height {
		return nums
	}
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	count := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			// res[(width*i+j)/c][(width*i+j)%c] = nums[i][j]
			res[count/c][count%c] = nums[i][j]
			count++
		}
	}
	return res
}

func Test_matrixReshape(t *testing.T) {
	a := [][]int{{1, 2}, {3, 4}}
	// a := [][]int{{1}}
	res := matrixReshape(a, 4, 1)
	t.Log(res)
	a = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	res = matrixReshape(a, 4, 3)
	t.Log(res)
}
