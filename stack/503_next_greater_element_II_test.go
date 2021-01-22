package stack

import (
	"fmt"
	"reflect"
	"testing"
)

//  503. Next Greater Element II
// 单调栈
func nextGreaterElementsIIA(nums []int) []int {
	length := len(nums)
	nums = append(nums, nums...)
	res := make([]int, 2*length)
	stack := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return res[:length]
}

// 不double切片，节省内存，但是取模运算降低了性能
func nextGreaterElementsII(nums []int) []int {
	length := len(nums)
	res := make([]int, len(nums))
	stack := []int{}
	for i := 2*len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i%length] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i%length] = -1
		} else {
			res[i%length] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i%length])
	}
	return res[:length]
}

func nextGreaterElements20210122(nums []int) []int {
	n := append(nums, nums...)
	fmt.Println(n)
	stack := make([]int, 0)
	res := make([]int, len(n))
	for i := len(n) - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= n[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}
		stack = append(stack, n[i])
		fmt.Println(stack)
		fmt.Println(res)
	}
	return res[:len(nums)]
}

func Test_nextGreaterElements20210122(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 2, 1}}, []int{2, -1, 2}},
		{"2", args{[]int{3, 2, 1}}, []int{-1, -1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements20210122(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements20210122() = %v, want %v", got, tt.want)
			}
		})
	}
}
