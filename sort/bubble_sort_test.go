package sort

import (
	"fmt"
	"testing"
)

func Bubble(nums []int) {
	// 两层循环，第一次外层循环将最大数冒泡至最右侧位置，
	// 第二次将次最大数冒泡至倒数第二位置
	// 时间复杂度 O(n2)
	// 空间复杂度 O(1)
	length := len(nums)
	var swapped int
	for i := length - 1; i >= 0; i-- {
		swapped = 0
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = 1
			}
		}
		if swapped == 0 {
			break
		}
	}
	return
}

func TestBubble(t *testing.T) {
	nums := []int{4,6,5,2,3,1}
	fmt.Println(nums)
	Bubble(nums)
	fmt.Println(nums)

}
