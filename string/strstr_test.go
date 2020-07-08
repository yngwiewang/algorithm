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

func strStr1(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	if needleLen > haystackLen {
		return -1
	}

	if needleLen == 0 {
		return 0
	}

	c := needle[0]
	for i := 0; i <= haystackLen-needleLen; i++ {
		if haystack[i] != c {
			continue
		}

		if haystack[i:i+needleLen] == needle {
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
