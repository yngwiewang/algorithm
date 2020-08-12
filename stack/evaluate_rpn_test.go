package stack

import (
	"strconv"
	"testing"
)

// 150. Evaluate Reverse Polish Notation

func evalRPN(tokens []string) int {
	var tmpStack []int
	operators := make(map[string]bool)
	for _, o := range [...]string{"+", "-", "*", "/"} {
		operators[o] = true
	}

	for _, v := range tokens {
		if !operators[v] {
			nu, _ := strconv.Atoi(v)
			tmpStack = append(tmpStack, nu)
		} else {
			b := tmpStack[(len(tmpStack) - 1)]
			a := tmpStack[(len(tmpStack) - 2)]
			tmpStack = tmpStack[:len(tmpStack)-2]
			var res int
			if v == "+" {
				res = a + b
			} else if v == "-" {
				res = a - b
			} else if v == "*" {
				res = a * b
			} else if v == "/" {
				res = a / b
			}
			tmpStack = append(tmpStack, res)
		}
	}
	return tmpStack[0]
}


// strconv.ParseInt is slower than strconv.Atoi
func evalRPN1(tokens []string) int {
	var stack []int
	for _, v := range tokens {
		switch v {
		case "+":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			res := a + b
			stack[len(stack)-2] = res
			stack = stack[:len(stack)-1]
		case "-":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			res := a - b
			stack[len(stack)-2] = res
			stack = stack[:len(stack)-1]
		case "*":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			res := a * b
			stack[len(stack)-2] = res
			stack = stack[:len(stack)-1]
		case "/":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			res := a / b
			stack[len(stack)-2] = res
			stack = stack[:len(stack)-1]
		default:
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

func Test_evalRPN1(t *testing.T) {
	s := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	t.Log(evalRPN(s))

}

func Benchmark_evalRPN1(b *testing.B) {
	s := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	for i := 0; i < b.N; i++ {
		evalRPN1(s)
	}
}
