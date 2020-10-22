package string

import (
	"fmt"
	"testing"
)

// [344] Reverse String

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(string(s))
}

func reverseString1(s []byte) {
	for i:= 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	fmt.Println(string(s))
}

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{"",args{[]byte("hello")}},
		{"",args{[]byte("")}},
		{"",args{[]byte("a")}},
		{"",args{[]byte("ab")}},
		{"",args{[]byte("abc")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString1(tt.args.s)
		})
	}
}
