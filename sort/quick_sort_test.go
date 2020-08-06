package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func QuickSort(n []int, p, r int) {
	if p >= r {
		return
	}
	q := partition(n, p, r)
	QuickSort(n, p, q-1)
	QuickSort(n, q+1, r)

}

func partition(n []int, p, r int) int {
	q := rand.Intn(r-p) + p
	pivot := n[q]
	n[q], n[r] = n[r], n[q]
	i := p
	for j := p; j < r; j++ {
		if n[j] < pivot {
			n[i], n[j] = n[j], n[i]
			i++
		}
	}
	n[i], n[r] = n[r], n[i]
	return i
}

func QuickSort2(n []int, p, r int) {
	if p < r {
		i, j := p, r
		q := rand.Intn(r-p) + p
		pivot := n[q]
		n[p], n[q] = n[q], n[p]
		for i < j {
			for i < j && n[j] >= pivot {
				j--
			}
			if i < j {
				n[i] = n[j]
				i++
			}
			for i < j && n[i] < pivot {
				i++
			}
			if i < j {
				n[j] = n[i]
			}
		}
		n[i] = pivot
		QuickSort2(n, p, i-1)
		QuickSort2(n, i+1, r)
	}
}

func QuickSort3(a []int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuickSort3(a[:left])
	QuickSort3(a[left+1:])

}

var n = RandSlice

func TestQuick(t *testing.T) {
	n := []int{2, 1, 4}
	QuickSort(n, 0, len(n)-1)
	fmt.Println(n)
}

func TestQuick2(t *testing.T) {
	n := []int{2, 1, 4}
	QuickSort2(n, 0, len(n)-1)
	fmt.Println(n)
}

//func TestQuick3(t *testing.T) {
//	for i:=0;i<100;i++{
//
//		n := RandomIntSlice(10000)
//		fmt.Println(n)
//		QuickSort3(n)
//		fmt.Println(n)
//	}
//}

func QC(n []int) {
	if len(n) < 2 {
		return
	}
	l, r := 0, len(n)-1
	q := rand.Intn(len(n))
	n[q], n[r] = n[r], n[q]
	for i := range n {
		if n[i] < n[r] {
			n[i], n[l] = n[l], n[i]
			l++
		}
	}
	n[l], n[r] = n[r], n[l]
	QC(n[:l])
	QC(n[l+1:])
}

func TestQc(t *testing.T) {
	n := []int{5, 7, 1, 2, 6, 4}
	QC(n)
	fmt.Println(n)
}

func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		fmt.Println(rand.Intn(10))
	}
	fmt.Println()
	for i := 0; i < 20; i++ {
		fmt.Println(rand.Int() % 10)
	}
}

func sort20200717(n []int, l, r int) {
	if l >= r {
		return
	}
	m := rand.Intn(r-l) + l
	pivot := n[m]
	n[r], n[m] = n[m], n[r]
	i := l
	for j := l; j < r; j++ {
		if n[j] <= pivot {
			n[i], n[j] = n[j], n[i]
			i++
		}
	}
	n[r], n[i] = n[i], n[r]
	sort20200717(n, l, i-1)
	sort20200717(n, i+1, r)
	return
}

func quickSort20200717(nums []int) []int {
	sort20200717(nums, 0, len(nums)-1)
	return nums
}

var a = []int{2, 1, 4, 6, 3, 5, 7, 1, 2, 3, 4, 5}

func TestQuickSort20200717(t *testing.T) {
	res := quickSort20200717(a)
	t.Log(res)
}

func sort20200717_1(n []int, l, r int) {
	if l >= r {
		return
	}
	i, j, p := l, r, n[l]
	for i <= j {
		if n[i] < p {
			i++
		} else if n[j] > p {
			j--
		} else {
			n[i], n[j] = n[j], n[i]
			i, j = i+1, j-1
		}
	}
	sort20200717_1(n, l, j)
	sort20200717_1(n, i, r)
}

func quickSort20200717_1(nums []int) []int {
	sort20200717_1(nums, 0, len(nums)-1)
	return nums
}

func TestQuickSort20200717_1(t *testing.T) {
	res := quickSort20200717_1(a)
	t.Log(res)
}

func quickSort20200806(n []int) {
	if len(n) < 2 {
		return
	}
	l, r := 0, len(n)-1
	p := rand.Intn(len(n))	
	n[p], n[r] = n[r], n[p]	
	for i := range n {
		if n[i] < n[r] {
			n[i], n[l] = n[l], n[i]
			l++
		}
	}
	n[l], n[r] = n[r], n[l]
	quickSort20200806(n[:l])
	quickSort20200806(n[l+1:])
}

func Test_quickSort20200806(t *testing.T) {
	n := []int{5, 7, 4, 2, 6, 4}
	quickSort20200806(n)
	t.Log(n)
}

func Test_slice(t *testing.T) {
	n := []int{5, 7, 1, 2, 6, 4}
	t.Log(n[0:0], len(n[0:0]))
}
