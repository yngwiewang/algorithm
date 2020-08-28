package string

import (
	"strings"
	"testing"
)

// 796. Rotate String
func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if A == "" && B == ""{
		return true
	}

	for i := 0; i < len(A); i++ {
		if A[i] == B[0] {
			if rotate(A, i) == B {
				return true
			}
		}
	}

	return false
}

func rotate(s string, n int) string {
	a := s[n:]
	b := s[:n]
	return a + b
}


func rotateString1(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if A == "" && B == ""{
		return true
	}
	newA := A + A
	return strings.Contains(newA, B)
}


func Test_rotateString(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"abcde", "cdeab"}, true},
		{"", args{"abcdefghijklmn", "nabcdefghijklm"}, true},
		{"", args{"abcde", "abced"}, false},
		{"", args{"abcde", " "}, false},
		{"", args{" ", "abced"}, false},
		{"", args{"abcde", ""}, false},
		{"", args{"", "abced"}, false},
		{"", args{"", ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateString1(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("rotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
