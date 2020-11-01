package stack

import (
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

func Test_nge2(t *testing.T) {
	a := []int{3, 2, 1}
	res := nextGreaterElementsII(a)
	t.Log(res)
}
