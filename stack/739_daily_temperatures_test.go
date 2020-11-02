package stack

import (
	"testing"
)

// 739. Daily Temperatures
// 单调栈
func dailyTemperaturesA(T []int) []int {
	res := make([]int, len(T))
	stack := [][2]int{}
	for i := len(T) - 1; i >= 0; i-- {
		tmp := 0
		for len(stack) > 0 && stack[len(stack)-1][0] <= T[i] {
			tmp += stack[len(stack)-1][1]
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = 0
			stack = append(stack, [2]int{T[i], 0})
		} else {
			res[i] = tmp + 1
			stack = append(stack, [2]int{T[i], res[i]}) // 此处用 res[i] 比 tmp + 1 性能更高
		}
	}
	return res
}

// 简化，不用数组做栈元素，性能没有提升
func dailyTemperaturesB(T []int) []int {
	res := make([]int, len(T))
	stack := []int{}
	for i := len(T) - 1; i >= 0; i-- {
		tmp := 0
		for j := i + 1; len(stack) > 0 && stack[len(stack)-1] <= T[i]; {
			tmp += res[j]
			j += res[j]
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 { // 如果先判断==0再else，性能降低
			res[i] = tmp + 1
		}

		stack = append(stack, T[i])
	}
	return res
}

// 记录位置而不是值
func dailyTemperaturesC(T []int) []int {
	res := make([]int, len(T))
	stack := []int{}
	for i := len(T) - 1; i >= 0; i-- {
		for len(stack) > 0 && T[stack[len(stack)-1]] <= T[i] {

			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return res
}

// 最快答案，正向循环，遇到比前面大的就修改其值
func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	stack := []int{}
	for i := range T {
		for len(stack) > 0 && T[stack[len(stack)-1]] < T[i] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

func Test_dayliTemperatures(t *testing.T) {
	// a := []int{73, 74, 75, 71, 69, 72, 76, 73}
	a := []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}
	res := dailyTemperatures(a)
	t.Log(res)
}
