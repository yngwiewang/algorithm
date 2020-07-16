package sort

import (
	"fmt"
	"testing"
)

func BubbleSort(nums []int) {
	// 两层循环，第一次外层循环将最大数冒泡至最右侧位置，
	// 第二次将次最大数冒泡至倒数第二位置
	// 时间复杂度 最好O(n)，最坏O(n2)，平均O(n2)
	// 空间复杂度 O(1)
	// 稳定
	length := len(nums)
	var swapped bool
	for i := length - 1; i >= 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return
}

func TestBubble(t *testing.T) {
	a := [][]int{{}, {1}, {2, -1}, {5, 7, 4, 6, 1, 2}, {7, 4, 5, 6, 2, 1}, {1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}}
	for _, l := range a {
		fmt.Println("origin list: ", l)
		res := BubbleSort20200716(l)
		fmt.Println("sorted list: ", res)
	}

}

func BubbleSort20200716(nums []int) []int {
	for i := len(nums) - 1; i > 0; i-- {
		swap := false
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
	return nums
}
