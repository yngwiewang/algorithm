package array

import (
	"reflect"
	"testing"
)

// 54. Spiral Matrix

const (
	right = iota
	down
	left
	up
)

// 四个方向螺旋遍历
func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	if row == 0 {
		return []int{}
	}
	col := len(matrix[0])
	if col == 0 {
		return []int{}
	}
	count := row * col
	trace := make([][2]int, 1, count)
	trace[0] = [2]int{0, -1}
	res := make([]int, 0, count)
	row--

	direction := right
	for len(trace) <= count {
		if direction == right {
			for i := 0; i < col; i++ {
				a, b := trace[len(trace)-1][0], trace[len(trace)-1][1]
				b++
				trace = append(trace, [2]int{a, b})
				res = append(res, matrix[a][b])
			}
			col--
			direction = down
			continue
		}
		if direction == down {
			for i := 0; i < row; i++ {
				a, b := trace[len(trace)-1][0], trace[len(trace)-1][1]
				a++
				trace = append(trace, [2]int{a, b})
				res = append(res, matrix[a][b])
			}
			row--
			direction = left
			continue
		}
		if direction == left {
			for i := 0; i < col; i++ {
				a, b := trace[len(trace)-1][0], trace[len(trace)-1][1]
				b--
				trace = append(trace, [2]int{a, b})
				res = append(res, matrix[a][b])
			}
			col--
			direction = up
			continue
		}
		if direction == up {
			for i := 0; i < row; i++ {
				a, b := trace[len(trace)-1][0], trace[len(trace)-1][1]
				a--
				trace = append(trace, [2]int{a, b})
				res = append(res, matrix[a][b])
			}
			row--
			direction = right
			continue
		}
	}
	// fmt.Println(trace)
	return res
}

// 简化代码，降低内存消耗
func spiralOrder1(matrix [][]int) []int {
	row := len(matrix)
	if row == 0 {
		return []int{}
	}
	col := len(matrix[0])
	if col == 0 {
		return []int{}
	}
	count := row * col
	a, b := 0, -1
	res := make([]int, 0, count)
	row--

	direction := right
	for len(res) < count {
		switch direction {
		case right:
			for i := 0; i < col; i++ {
				b++
				res = append(res, matrix[a][b])
			}
			col--
			direction = down
		case down:
			for i := 0; i < row; i++ {
				a++
				res = append(res, matrix[a][b])
			}
			row--
			direction = left
		case left:
			for i := 0; i < col; i++ {
				b--
				res = append(res, matrix[a][b])
			}
			col--
			direction = up
		case up:
			for i := 0; i < row; i++ {
				a--
				res = append(res, matrix[a][b])
			}
			row--
			direction = right
		}
	}
	return res
}

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{"", args{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {1, 2, 3, 4}}}, []int{1, 2, 3, 4, 8, 12, 4, 3, 2, 1, 9, 5, 6, 7, 11, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder1(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
