package array

import (
	"testing"
)

// 75. Sort Colors

// 用一个数组存储三个的个数，重写原slice的时候用循环的话消耗更多一点内存
func sortColors(nums []int) {
	colors := [3]int{}
	for _, v := range nums {
		colors[v]++
	}
	c := 0
	for j := 0; j < 3; j++ {
		for i := 0; i < colors[j]; i++ {
			nums[c] = j
			c++
		}
	}

	// for i := 0; i < colors[1]; i++ {
	// 	nums[c] = 1
	// 	c++
	// }
	// for i := 0; i < colors[2]; i++ {
	// 	nums[c] = 2
	// 	c++
	// }
	return
}

// 三指针：0，2，当前,原地排序
func sortColors1(nums []int) {
	l, c, r := 0, 0, len(nums)-1
	for c <= r {
		switch nums[c] {
		case 0:
			nums[l], nums[c] = nums[c], nums[l]
			c, l = c+1, l+1
		case 2:
			nums[r], nums[c] = nums[c], nums[r]
			r--
		default:
			c++
		}
	}
}
func Test_sortColors(t *testing.T) {
	c := []int{2, 0, 1}
	// c := []int{2, 0, 2, 1, 1, 0}
	// c := []int{2}
	sortColors1(c)
	t.Log(c)
}
