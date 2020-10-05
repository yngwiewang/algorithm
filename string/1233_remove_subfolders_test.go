package string

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

// 1233. Remove Sub-Folders from the Filesystem
// 先排序再判断后一个字符串是否以前一个字符串加"/"为前缀
// 注意排序用 sort.Strings 函数
func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	fmt.Println(folder)
	res := []string{folder[0]}
	j := 0
	for i := 1; i < len(folder); i++ {
		if !strings.HasPrefix(folder[i], res[j]+"/") {
			res = append(res, folder[i])
			j++
		}
	}
	return res
}

func Test_removeSubFolders(t *testing.T) {
	// a := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	a := []string{"/a", "/a/b/c", "/a/b/d"}
	// a := []string{"/a/b/c", "/a/b/ca", "/a/b/d"}
	// a := []string{"/ah/al/am", "/ah/al"}
	res := removeSubfolders(a)
	t.Log(res)
}
