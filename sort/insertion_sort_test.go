package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func RandomIntSlice() (r []int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100000; i++ {
		r = append(r, rand.Intn(1000000))
	}
	return
}

var RandSlice = RandomIntSlice()

func InsertionSort(nums []int) {
	// 取未排序区间中的元素，在已排序区间中找到合适的位置将其插入
	// 时间复杂度O(n2)
	// 空间复杂度O(1)
	// 稳定
	length := len(nums)
	for i := 1; i < length; i++ {
		current := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if current < nums[j] {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = current
	}
}

func TestInsertionSort(t *testing.T) {
	nums := []int{4, 6, 5, 2, 3, 1}
	fmt.Println("before:", nums)
	InsertionSort(nums)
	fmt.Println("after:", nums)
}

func BenchmarkBubbleInsertion1(b *testing.B) {
	nums := RandSlice
	for i := 0; i < b.N; i++ {
		BubbleSort(nums)
	}
}

func BenchmarkBubbleInsertion2(b *testing.B) {
	nums := RandSlice
	for i := 0; i < b.N; i++ {
		InsertionSort(nums)
	}
}

// 插入比冒泡交换元素次数少，性能更好
//BenchmarkBubbleInsertion1-8            1        11977649400 ns/op
//BenchmarkBubbleInsertion2-8        15739             74228 ns/op
