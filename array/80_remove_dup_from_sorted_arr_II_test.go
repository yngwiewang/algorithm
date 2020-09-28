package array

import "testing"

// 80. Remove Duplicates from Sorted Array II

func removeDuplicatesII(nums []int) int {
	res := len(nums)
	for i := 0; i < len(nums); {
		v := nums[i]
		if i < res-1 && nums[i] == nums[i+1] {
			i++
			// 此处优化为记录t后copy一次，可以100%
			j, t := i+1, 0
			for ; j < res && nums[j] == v; j++ {
				t++
			}
			if t > 0 {
				copy(nums[i+1:], nums[j:])
				res -= t
			}
			i++
		} else {
			i++
		}
	}
	return res
}

func Test_removeDuplicatesII(t *testing.T) {
	a := []int{0, 0, 1, 1, 2, 2, 2, 3}
	res := removeDuplicatesII(a)
	t.Log(a)
	t.Log(res)
}
