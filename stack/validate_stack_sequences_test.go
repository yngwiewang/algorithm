package stack

import "testing"

// 946. Validate Stack Sequences
func validateStackSequences1(pushed []int, popped []int) bool {
	var stack []int
	var count int
	for _, v := range pushed {
		stack = append(stack, v)
		for len(popped) > 0 && len(stack) > 0 && popped[0] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
			count++
		}
	}
	return count == len(pushed)
}

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	var j int
	for _, v := range pushed {
		stack = append(stack, v)
		for i := 0; len(stack) > 0 && popped[j] == stack[len(stack)-1] && i < len(popped); i++ {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return j == len(pushed)
}

func Test_validateStackSequences(t *testing.T) {
	type args struct {
		pushed []int
		popped []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1: ", args{[]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}}, true},
		{"2: ", args{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}}, false},
		{"3: ", args{[]int{1, 2, 3, 4, 5}, []int{3, 2, 1, 4, 5}}, true},
		{"4: ", args{[]int{1, 2, 3, 4, 5}, []int{3, 2, 4, 1, 5}}, true},
		{"5: ", args{[]int{1, 2, 3, 4, 5}, []int{3, 1, 4, 5, 2}}, false},
		{"6: ", args{[]int{1, 2, 3}, []int{3, 1, 2}}, true},
		{"7: ", args{[]int{1, 2, 3}, []int{3, 2, 1}}, true},
		{"8: ", args{[]int{1, 2, 3}, []int{1, 3, 2}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateStackSequences(tt.args.pushed, tt.args.popped); got != tt.want {
				t.Errorf("validateStackSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
