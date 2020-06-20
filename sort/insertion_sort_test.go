package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func RandomIntSlice(n int) []int {
	rand.Seed(time.Now().UnixNano())
	r := make([]int, n, n)
	for i := 0; i < n; i++ {
		r[i] = rand.Intn(1000)
	}
	return r
}

var RandSlice = RandomIntSlice(100000)

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

//func BenchmarkBubbleSort(b *testing.B) {
//	nums := RandSlice
//	for i := 0; i < b.N; i++ {
//		BubbleSort(nums)
//	}
//}

//func BenchmarkInsertionSort(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		nums := RandSlice
//		InsertionSort(nums)
//	}
//}

// 插入比冒泡交换元素次数少，性能更好
//BenchmarkBubbleInsertion1-8            1        11977649400 ns/op
//BenchmarkBubbleInsertion2-8        15739             74228 ns/op

//func BenchmarkMergeSort(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		nums := RandSlice
//		MergeSort(nums, 0, len(nums)-1)
//	}
//}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//nums:=RandomIntSlice(100000)
		//b.StartTimer()
		QuickSort(RandSlice, 0, len(RandSlice)-1)
		//b.StopTimer()
	}
}

func BenchmarkQuickSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//nums:=RandomIntSlice(100000)
		//b.StartTimer()
		QuickSort2(RandSlice, 0, len(RandSlice)-1)
		//b.StopTimer()
	}
}

func BenchmarkQuickSort3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//nums:=RandomIntSlice(100000)
		//b.StartTimer()
		QuickSort3(RandSlice)
		//b.StopTimer()
	}
}

func BenchmarkQuickSort4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//nums:=RandomIntSlice(100000)
		//b.StartTimer()
		QC(RandSlice)
		//b.StopTimer()
	}
}

func BenchmarkRand1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		a := rand.Intn(100)
		a = a >> 1
	}
}

func BenchmarkRand2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		a := rand.Int() % 100
		a = a >> 1
	}
}

//func TestQuick2(t *testing.T){
//	nums := RandSlice
//	QuickSort(nums,0,len(nums)-1)
//	fmt.Println(nums)
//	nums = RandSlice
//	QuickSort1(nums,0,len(nums)-1)
//	fmt.Println(nums)
//
//}
