package array

import (
	"fmt"
	"testing"
)

// 48. Rotate Image

// 计算要旋转的圈数（层数），再计算每一层要交换的次数，逐个交换
func rotateImage(matrix [][]int) {
	circle := len(matrix) / 2
	for r := 0; r < circle; r++ {
		head, tail := r, len(matrix)-r-1
		swapCounts := len(matrix) - r*2 - 1
		for i := 0; i < swapCounts; i++ {
			tmp := matrix[head][head+i]
			matrix[head][head+i] = matrix[tail-i][head]
			matrix[tail-i][head] = matrix[tail][tail-i]
			matrix[tail][tail-i] = matrix[head+i][tail]
			matrix[head+i][tail] = tmp
		}
	}
	fmt.Println(matrix)
}

// 先对角线镜像交换，再左右镜像交换
func rotateImage1(matrix [][]int) {
	for _, v := range matrix {
		fmt.Println(v)
	}
	fmt.Println("-------------")
	edge := len(matrix)
	// “左下-右上”对角线镜像交换
	for i := 0; i <= edge; i++ {
		for j := i + 1; j < edge; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for _, v := range matrix {
		fmt.Println(v)
	}
	fmt.Println("-------------")
	// “左右水平”镜像交换
	for i := 0; i < edge; i++ {
		for j := 0; j < edge/2; j++ {
			matrix[i][j], matrix[i][edge-j-1] = matrix[i][edge-j-1], matrix[i][j]
		}
	}
	for _, v := range matrix {
		fmt.Println(v)
	}
}

func Test_rotateImage(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{"1-9", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}},
		{"1-16", args{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotateImage1(tt.args.matrix)
		})
	}
}
