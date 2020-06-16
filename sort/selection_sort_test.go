package sort

import (
	"fmt"
	"testing"
)

func SelectionSort(nums []int) {
	// 时间复杂度O(n2)
	// 空间复杂度O(1)
	// 不稳定
	length := len(nums)
	for i := 0; i < length; i++ {
		min := nums[i]
		minPos := i
		j := i + 1
		for ; j < length; j++ {
			if nums[j] < min {
				min = nums[j]
				minPos = j
			}
		}
		nums[i], nums[minPos] = nums[minPos], nums[i]
	}
}

func TestSelectionSort(t *testing.T) {
	nums := []int{4, 6, 5, 2, 3, 1}
	fmt.Println("before:", nums)
	SelectionSort(nums)
	fmt.Println("after:", nums)
}
