package sort

import (
	"fmt"
	"testing"
)

var tmp = make([]int, 200000)

func Merge(nums []int, low, mid, high int) {
	i, j, size := low, mid+1, 0
	for ; i <= mid && j <= high; size++ {
		if nums[i] <= nums[j] {
			tmp[size] = nums[i]
			i++
		} else {
			tmp[size] = nums[j]
			j++
		}
	}
	for i <= mid {
		tmp[size] = nums[i]
		size++
		i++
	}
	for j <= high {
		tmp[size] = nums[j]
		size++
		j++
	}
	for i := 0; i < size; i++ {
		nums[low+i] = tmp[i]
	}
}

func MergeSort(nums []int, low, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	MergeSort(nums, low, mid)
	MergeSort(nums, mid+1, high)
	Merge(nums, low, mid, high)
}

func TestMergeSort(t *testing.T) {
	a := []int{4, 6, 5, 3, 1, 2, 54, 6, 435, 32, 5, 345, 34, 534, 534, 534, 7}
	MergeSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func mergeSort20200716(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort20200716(nums[:mid])
	right := mergeSort20200716(nums[mid:])
	return merge20200716(left, right)
}

func merge20200716(l, r []int) []int {
	res := make([]int, len(l)+len(r))
	i, j, k := 0, 0, 0
	for i < len(l) && j < len(r) {
		if l[i] <= r[j] {
			res[k] = l[i]
			i++
		} else {
			res[k] = r[j]
			j++
		}
		k++
	}
	// 加判断性能更好
	if i < len(l) {
		copy(res[k:], l[i:])
	} else {
		copy(res[k:], r[j:])
	}
	return res
}

func TestMergeSort20200716(t *testing.T) {
	a := []int{2, -1, 4, 3, 9, 7, 8, 6}
	res := mergeSort20200716(a)
	t.Log(res)
}
