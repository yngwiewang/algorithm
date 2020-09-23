package array

import (
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 57. Insert Interval

func insertInterval(intervals [][]int, newInterval []int) [][]int {

	i := 0
	for ; i < len(intervals) && intervals[i][1] < newInterval[0]; i++ {

	}
	j := i
	for ; i < len(intervals) && newInterval[1] >= intervals[i][0]; i++ {
		newInterval = []int{common.MinInt(newInterval[0], intervals[i][0]),
			common.MaxInt(newInterval[1], intervals[i][1])}
	}

	right := make([][]int, len(intervals[i:]))
	copy(right, intervals[i:])

	return append(append(intervals[:j], newInterval), right...)
}

func Test_insertInterval(t *testing.T) {
	// a := [][]int{{1, 3}, {6, 9}}
	// b := []int{2, 5}
	// a := [][]int{}
	// b := []int{2, 5}
	// a := [][]int{{1,3}}
	// b := []int{2, 5}
	a := [][]int{{1, 5}}
	b := []int{0, 0}
	t.Log(insertInterval(a, b))
}
