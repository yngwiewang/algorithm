package array

import (
	"reflect"
	"testing"
)

// 977. Squares of a Sorted Array
func sortedSquares(a []int) []int {
	res := make([]int, len(a))
	k := len(a) - 1
	i, j := 0, len(a)-1
	for i <= j {
		if a[i]*a[i] > a[j]*a[j] {
			res[k] = a[i] * a[i]
			i++
		} else {
			res[k] = a[j] * a[j]
			j--
		}
		k--
	}
	return res
}

func Test_sortedSquares(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"-7,-3,2,3,11", args{[]int{-7, -3, 2, 3, 11}}, []int{4, 9, 9, 49, 121}},
		{"-4,-1,0,-3,10", args{[]int{-4, -1, 0, -3, 10}}, []int{0, 1, 9, 16, 100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
