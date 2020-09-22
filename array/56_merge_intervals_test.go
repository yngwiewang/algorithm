package array

import (
	"sort"
	"testing"
)

// 56. Merge Intervals

type sortSlice2 [][]int

func (s sortSlice2) Len() int { return len(s) }

func (s sortSlice2) Less(i, j int) bool {
	if s[i][0] != s[j][0] {
		return s[i][0] < s[j][0]
	}
	return s[i][1] < s[j][1]
}

func (s sortSlice2) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// 先根据第一个元素排序，再处理区间包含
// 由于先填充第一个元素到 res 中所以要先判空
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	s := sortSlice2(intervals)
	sort.Sort(s)
	res := [][]int{s[0]}

	for i := 1; i < len(s); i++ {
		pre := res[len(res)-1]
		if s[i][0] <= pre[1] {
			if pre[1] < s[i][1] {
				pre[1] = s[i][1]
			}
			continue
		}
		res = append(res, s[i])

	}
	return res
}

// 简化排序代码
// ! 由于有包含的判断，所以无需比较第二个元素
// ! sort.Slice 比上面的 sort.Sort 多两次内存分配，性能下降
func merge1(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		pre := res[len(res)-1]
		if intervals[i][0] <= pre[1] {
			if pre[1] < intervals[i][1] {
				pre[1] = intervals[i][1]
			}
			continue
		}
		res = append(res, intervals[i])

	}
	return res
}

func Test_mergeIntervals(t *testing.T) {
	intervals := [][]int{{1, 4}, {2, 3}}
	// intervals := [][]int{{1, 3}, {8, 10}, {2, 6}, {15, 18}}
	// intervals := [][]int{}
	res := merge(intervals)
	t.Log(res)
}

func Benchmark_mergeIntervals(b *testing.B) {
	intervals := [][]int{{1, 3}, {8, 10}, {2, 6}, {7, 19}}
	for i := 0; i < b.N; i++ {
		merge(intervals)
	}
}
