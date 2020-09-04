package string

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

// 49. Group Anagrams
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		b := []byte(s)
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		m[string(b)] = append(m[string(b)], s)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// * 内存更少
func groupAnagrams1(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		count := make([]int, 26)
		for i := 0; i < len(s); i++ {
			count[s[i]-'a']++
		}
		var sb strings.Builder
		for _, v := range count {
			sb.WriteByte(byte(v))
			sb.WriteByte('#')
		}
		m[sb.String()] = append(m[sb.String()], s)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// leetcode答案，最快
func groupAnagrams2(strs []string) [][]string {
	m := make(map[[26]byte]int)
	res := make([][]string, len(m))
	for _, s := range strs {
		count := [26]byte{}
		for i := 0; i < len(s); i++ {
			count[s[i]-'a']++
			fmt.Println(s[i]-'a')
		}
		if idx, found := m[count]; found {
			res[idx] = append(res[idx], s)
		} else {
			res = append(res, []string{s})
			m[count] = len(res) - 1
		}
	}
	return res
}

func Benchmark_groupAnagrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
		_ = groupAnagrams2(a)
	}
}

func Test_groupAnagrams(t *testing.T) {
	a := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	b := []string{""}
	c := []string{"a"}
	res := groupAnagrams2(a)
	fmt.Println(res)
	res = groupAnagrams2(b)
	fmt.Println(res)
	res = groupAnagrams2(c)
	fmt.Println(res)
}
