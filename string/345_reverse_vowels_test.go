package string

import (
	"testing"
)

//345. Reverse Vowels of a String

func reverseVowels(s string) string {
	bs := []byte(s)
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	i, j := 0, len(bs)-1
	for i < j {
		if _, ok := vowels[bs[i]]; !ok {
			i++
		} else if _, ok := vowels[bs[j]]; !ok {
			j--
		} else {
			bs[i], bs[j] = bs[j], bs[i]
			i, j = i+1, j-1
		}
	}
	return string(bs)
}

func Test_reverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"aeiou", args{"aeiou"}, "uoiea"},
		{"hello", args{"hello"}, "holle"},
		{"leetcode", args{"leetcode"}, "leotcede"},
		{"", args{""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
