package array

import (
	"reflect"
	"testing"
)

//66. Plus One
func plusOne(digits []int) []int {
	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			digits[i]++
		} else {
			digits[i] += carry
		}
		if digits[i] < 10 {
			return digits
		}
		digits[i] = digits[i] % 10
		carry = 1
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

// leetcode内存占用最小的答案，代码要简洁很多
// 重点是是要没有因为digits[i]<9而return，就要一直进位
func plusOne1(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	digits = append([]int{1}, digits...)
	return digits
}

func BenchmarkPlusOne(b *testing.B) {
	a := make([]int, 999999, 999999)
	for i := 0; i < 999999; i++ {
		a[i] = 9
	}
	for i := 0; i < b.N; i++ {
		_ = plusOne1(a)
	}
}

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{9}}, []int{1, 0}},
		{"", args{[]int{0}}, []int{1}},
		{"", args{[]int{1, 2, 3}}, []int{1, 2, 4}},
		{"", args{[]int{1, 2, 9}}, []int{1, 3, 0}},
		{"", args{[]int{1, 9, 9}}, []int{2, 0, 0}},
		{"", args{[]int{9, 9, 9}}, []int{1, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne1(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
