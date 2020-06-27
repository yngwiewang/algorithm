package array

import (
	"fmt"
	"testing"
)

// leetcode 283

func moveZeroes(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
			i--
			n--
		}
	}
}

func moveZeroes1(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			copy(nums[i:], nums[i+1:])
			nums[len(nums)-1] = 0
			i--
			n--
		}
	}
}

func moveZeroes2(nums []int) {
	cur := 0
	for _, v := range nums {
		if v != 0 {
			nums[cur] = v
			cur++
		}
	}
	for ; cur < len(nums); cur++ {
		nums[cur] = 0
	}
}

func moveZeroes3(nums []int) {
	j := 0
	for i, v := range nums {
		if v != 0 {
			nums[j] = v
			if i != j {
				nums[i] = 0
			}
			j++
		}
	}

}

func moveZeroes4(nums []int) {
	j := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

func BenchmarkMoveZeros(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []int{0, 0, 1, 0, 0, 0, 3, 12, 0, 0, 0}
		moveZeroes(a)
	}
}

func BenchmarkMoveZeros1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []int{0, 0, 1, 0, 0, 0, 3, 12, 0, 0, 0}
		moveZeroes1(a)
	}
}

func TestMoveZeros(t *testing.T) {
	//a := []int{0, 0, 1, 0, 0, 0, 3, 12, 0, 0, 0,}
	a := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0, 0}
	moveZeroes4(a)
	fmt.Println(a)
}
