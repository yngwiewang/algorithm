package string

import (
	"fmt"
	"testing"
)

// leetcode 28

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0
		for j < len(needle) {
			if haystack[i+j] != needle[j] {
				break
			}
			j++
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}

func TestStrStr(t *testing.T) {
	a := "a"
	b := "a"
	fmt.Println(strStr(a, b))
}
