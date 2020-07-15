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
	n := []int{5, 7, 1, 4, 3, 2}
	QuickSort(n, 0, len(n)-1)
	fmt.Println(n)
}

func TestQuick2(t *testing.T) {
	n := []int{5, 7, 1, 4, 3, 2}
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
	n := []int{5, 7, 1, 4, 3, 2}
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

func QuickSort4(n []int, p, r int) {
	if p >= r {
		return
	}
	q := partition4(n, p, r)
	QuickSort4(n, p, q-1)
	QuickSort4(n, q+1, r)

}

func partition4(n []int, p, r int) int {
	q := rand.Intn(r-p) + p
	pivot := n[q]

	for {
		for n[p] <= pivot {
			p++
		}
		for n[r] > pivot {
			r--
		}
		if p >= r {
			break
		}
		n[p], n[r] = n[r], n[p]
	}
	return r
}

func TestQuickSort4(t *testing.T) {
	n := []int{5, 1, 1, 2, 0, 0}
	QuickSort4(n, 0, len(n)-1)
	fmt.Println(n)
}
