package string

import (
	"strings"
	"testing"
)

// 71. Simplify Path

func simplifyPath(path string) string {
	words := strings.Split(path, "/")
	ss := []string{}
	for _, w := range words {
		switch w {
		case ".", "":

		case "..":
			if len(ss) > 0 {
				ss = ss[:len(ss)-1]
			}
		default:
			ss = append(ss, w)
		}
	}
	res := "/" + strings.Join(ss, "/")
	return res
}

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"/.../.....", args{"/.../....."}, "/.../....."},
		{"/.../.", args{"/.../."}, "/..."},
		{"/.../..", args{"/.../.."}, "/"},
		{"/../", args{"/../"}, "/"},
		{"/home/", args{"/home/"}, "/home"},
		{"/a/./b/../../c/", args{"/a/./b/../../c/"}, "/c"},
		{"/a/../../b/../c//.//", args{"/a/../../b/../c//.//"}, "/c"},
		{"/a//b////c/d//././/..", args{"/a//b////c/d//././/.."}, "/a/b/c"},
		{"/.../..", args{"/.../.."}, "/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
