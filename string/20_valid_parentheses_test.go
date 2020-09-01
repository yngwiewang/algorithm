package string

import "testing"

// 20. Valid Parentheses

func isValid(s string) bool {
	stack := []rune{}
	for _, r := range s {
		switch r {
		case '{':
			stack = append(stack, r)
		case '(':
			stack = append(stack, r)
		case '[':
			stack = append(stack, r)
		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

// byte更快，内存消耗更少
func isValid1(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case byte('{'):
			stack = append(stack, s[i])
		case byte('('):
			stack = append(stack, s[i])
		case byte('['):
			stack = append(stack, s[i])
		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == byte('{') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == byte('[') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == byte('(') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

// 代码稍微整洁一点但是性能不如isValid1
func isValid2(s string) bool {
	stack := []byte{}
	for _, r := range s {
		if r == '{' || r == '[' || r == '(' {
			stack = append(stack, byte(r))
		} else {
			if len(stack) == 0 {
				return false
			}
			switch r {
			case '}':
				if stack[len(stack)-1] != byte('{') {
					return false
				}
				stack = stack[:len(stack)-1]
			case ']':
				if stack[len(stack)-1] != byte('[') {
					return false
				}
				stack = stack[:len(stack)-1]
			case ')':
				if stack[len(stack)-1] != byte('(') {
					return false
				}
				stack = stack[:len(stack)-1]

			}
		}
	}
	return len(stack) == 0
}

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"]", args{"]"}, false},
		{"{", args{"{"}, false},
		{"{[}]", args{"{[}]"}, false},
		{"()[]{}", args{"()[]{}"}, true},
		{"(]", args{"(]"}, false},
		{"{[]}", args{"{[]}"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid2(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_isValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := "{([{)[})[){"
		_ = isValid2(a)
	}
}
