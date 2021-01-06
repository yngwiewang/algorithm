package dict

import (
	"testing"
)

// 316. Remove Duplicate Letters

func removeDuplicateLetters(s string) string {
	count := [26]int{}
	for _, c := range s {
		count[c-'a']++
	}
	visited := [26]bool{}
	stack := make([]rune, 0)
	for _, c := range s {
		count[c-'a']--
		if visited[c-'a'] {
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > c && count[stack[len(stack)-1]-'a'] > 0 {
			visited[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, c)
		visited[c-'a'] = true
	}
	return string(stack)
}

func Test_removeDuplicateLetters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"acbabc", args{"acbabc"}, "abc"},
		{"abcacb", args{"abcacb"}, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateLetters(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicateLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
