package array

import (
	"testing"
)

// 88. Merge Sorted Array

// n2里面连续相同的数一次性插入n1
func mergeSortedArray1(n1 []int, m int, n2 []int, n int) {
	i, j, cnt := 0, 0, 0
	for {
		if cnt == n {
			return
		}
		if i >= m+cnt {
			copy(n1[i:], n2[j:])
			return
		}
		t := 0
		jt := j
		for j < n && n2[j] <= n1[i] {
			t, j = t+1, j+1
		}
		if t > 0 {
			copy(n1[i+t:], n1[i:])
			for a := i; a < i+t; a++ {
				n1[a] = n2[jt]
				jt++
			}
			i += t
			cnt += t
			continue
		}
		i++
	}
}

// 从后往前，从大往小插入
func mergeSortedArray(n1 []int, m int, n2 []int, n int) {
	if m == 0 {
		copy(n1, n2)
		return
	}
	l := m + n - 1
	m, n = m-1, n-1
	for ; m >= 0 && n >= 0; l-- {
		if n1[m] > n2[n] {
			n1[l] = n1[m]
			m--
		} else {
			n1[l] = n2[n]
			n--
		}
	}
	for ; n >= 0; l, n = l-1, n-1 {
		n1[l] = n2[n]
	}
}

func Test_mergeSortedArray(t *testing.T) {
	n1 := []int{2, 0}
	m := 1
	n2 := []int{1}
	n := 1

	mergeSortedArray(n1, m, n2, n)
	t.Log(n1)
}
