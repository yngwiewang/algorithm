package string

import "testing"

//925. Long Pressed Name

func isLongPressedName(name string, typed string) bool {
	if len(name) > len(typed) || name[0] != typed[0] || name[len(name)-1] != typed[len(typed)-1] {
		return false
	}
	i, j := 0, 0
	for j < len(typed) {
		if i < len(name) && name[i] == typed[j] {
			i, j = i+1, j+1
		} else {
			if typed[j] == typed[j-1] {
				j++
				continue
			}
			return false
		}
	}
	return true
}

func Test_isLongPressedName(t *testing.T) {
	type args struct {
		name  string
		typed string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"pyplrz", args{"pyplrz", "ppypllr"}, false},
		{"alex", args{"alex", "aaleex"}, true},
		{"saeed", args{"saeed", "ssaaedd"}, false},
		{"leelee", args{"leelee", "lleeelee"}, true},
		{"laiden", args{"laiden", "laiden"}, true},
		{"aa", args{"a", "a"}, true},
		{"aa", args{"a", "aa"}, true},
		{"aa", args{"aa", "a"}, false},
		{"a", args{"a", "ba"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLongPressedName(tt.args.name, tt.args.typed); got != tt.want {
				t.Errorf("isLongPressedName() = %v, want %v", got, tt.want)
			}
		})
	}
}
