package array

import (
	"reflect"
	"testing"
)

// 167. Two Sum II - Input array is sorted

func twoSumInSortedArray(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] == target {
			break
		}
		if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return []int{i + 1, j + 1}
}

func Test_twoSumInSortedArray(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"2,7,11,15", args{[]int{2, 7, 11, 15}, 9}, []int{1, 2}},
		{"2,3,4", args{[]int{2, 3, 4}, 6}, []int{1, 3}},
		{"-1,0", args{[]int{-1, 0}, -1}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumInSortedArray(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
