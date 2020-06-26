package array

import (
	"fmt"
	"testing"
)

// 未实现
//func removeDuplicates(nums []int) int {
//	n := len(nums)
//	//x := cap(nums)
//	for i := 0; i < len(nums); i++ {
//		j := i + 1
//		for ; j < n+1 && nums[j] == nums[i]; j++ {
//			n--
//		}
//		copy(nums[i+1:] ,nums[j:])
//		if j >= n{
//			break
//		}
//	}
//	return n
//}

// 性能低
//func removeDuplicates(nums []int) int{
//
//	for i:=len(nums)-1;i>0;i--{
//		if nums[i] == nums[i-1]{
//			nums = append(nums[:i], nums[i+1:]...)
//		}
//	}
//	return len(nums)
//}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	i, j := 0, 1
	for ; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			if i != j { // 可以去掉，但去掉后性能下降
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return i + 1
}

func TestRemoveDup(t *testing.T) {
	a := []int{1, 1, 1, 1}
	x := removeDuplicates(a)
	fmt.Println(a)
	fmt.Println(x)
	a = []int{1, 1, 1}
	x = removeDuplicates(a)
	fmt.Println(a)
	fmt.Println(x)
	a = []int{1, 1}
	x = removeDuplicates(a)
	fmt.Println(a)
	fmt.Println(x)
	a = []int{1}
	x = removeDuplicates(a)
	fmt.Println(a)
	fmt.Println(x)
	a = []int{}
	x = removeDuplicates(a)
	fmt.Println(a)
	fmt.Println(x)
	a = []int{1, 2, 2}
	x = removeDuplicates(a)
	fmt.Println(a)
	fmt.Println(x)
	a = []int{1, 1, 2}
	x = removeDuplicates(a)
	fmt.Println(a)
	fmt.Println(x)
	a = []int{1, 1, 2, 2}
	x = removeDuplicates(a)
	fmt.Println(a)
	fmt.Println(x)
	a = []int{1, 1, 2, 2, 3, 3, 4, 4, 5}
	x = removeDuplicates(a)
	fmt.Println(a)
	fmt.Println(x)
	a = []int{-3, -1, -1, 0, 0, 0, 0, 0, 2}
	x = removeDuplicates(a)
	fmt.Println(a)
	fmt.Println(x)
}
