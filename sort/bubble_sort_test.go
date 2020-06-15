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
	nums := []int{4, 6, 5, 2, 3, 1}
	fmt.Println(nums)
	BubbleSort(nums)
	fmt.Println(nums)

}
