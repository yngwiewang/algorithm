package string

import "testing"

// 1003. Check If Word Is Valid After Substitutions

// 使用栈
func isValidWord(s string) bool {
	stack := make([]byte, 0, 3)
	for i := 0; i < len(s); {
		if len(s)-i > 1 && len(stack) > 0 && s[i:i+2] == "bc" && stack[len(stack)-1] == 'a' {
			stack = stack[:len(stack)-1]
			i += 2
			continue
		}
		if len(stack) > 1 && s[i] == 'c' && stack[len(stack)-2] == 'a' && stack[len(stack)-1] == 'b' {
			stack = stack[:len(stack)-2]
			i++
			continue
		}
		stack = append(stack, s[i])
		i++
	}
	return len(stack) == 0
}

// leetcode最快答案，简化判断
func isValidWord1(s string) bool {
	stack := make([]byte, 0, 3)
	for i := range s {
		if s[i] == 'c' {
			if len(stack) == 0 || stack[len(stack)-1] != 'b' {
				return false
			}
			stack = stack[:len(stack)-1]
			if len(stack) == 0 || stack[len(stack)-1] != 'a' {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return len(stack) == 0
}

func Test_isValidWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"a", args{"a"}, false},
		{"ab", args{"ab"}, false},
		{"abc", args{"abc"}, true},
		{"aabcbc", args{"aabcbc"}, true},
		{"aabcb", args{"aabcb"}, false},
		{"abccba", args{"abccba"}, false},
		{"cababc", args{"cababc"}, false},
		{"abcabcababcc", args{"abcabcababcc"}, true},
		{"ababccaabcbcababcabcc", args{"ababccaabcbcababcabcc"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidWord1(tt.args.s); got != tt.want {
				t.Errorf("isValidWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
