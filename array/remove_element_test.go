package array

import (
	"testing"
)

// 27. Remove Element
// 第一遍自己写的暴力指针法，程序比较乱
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for i < len(nums) {
		for nums[i] != val {
			i++
			if i == len(nums)-1 {
				if nums[i] == val {
					return i
				}
				return i + 1
			}
		}
		j := i + 1
		for ; nums[j] == val; j++ {
			if j == len(nums)-1 {
				if nums[j] == val {
					nums[i], nums[j] = nums[j], nums[i]
				}
				return i
			}
		}
		nums[i], nums[j] = nums[j], nums[i]

	}
	return i
}

// 小浩算法的append实现，比较简洁
func removeElement1(nums []int, val int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return i
}

// ! 自己写的头尾双指针，简洁一些，返回时注意判断i的位置是不是val，
// ! 还要判断空slice，
// ! 速度最快，内存占用最低
func removeElement2(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i] != val {
			i++
		} else if nums[j] == val {
			j--
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			i, j = i+1, j-1
		}
	}
	if nums[i] == val {
		return i
	}
	return i + 1
}

// https://books.halfrost.com/leetcode/ChapterFour/0027.Remove-Element/
func removeElement3(nums []int, val int) int {
	j := 0
	for i := range nums {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
				j++
			} else {
				j++
			}
		}
	}
	return j
}

// leetcode 内存最小的答案，比cookbook更加简洁
func removeElement4(nums []int, val int) int {
	last := 0
	for i := range nums {
		if nums[i] != val {
			nums[last] = nums[i]
			last++
		}
	}
	return last
}

// 复习一遍，不用交换，只要更新nums[cur]即可
func removeElement5(nums []int, val int) int {
	cur := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[cur], nums[i] = nums[i], nums[cur]
			cur++
		}
	}
	return cur
}

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{}, 1}, 0},
		{"2", args{[]int{2}, 1}, 1},
		{"2,2", args{[]int{2, 2}, 2}, 0},
		{"2,1", args{[]int{2, 1}, 1}, 1},
		{"2", args{[]int{2}, 2}, 0},
		{"", args{[]int{2, 2, 2}, 1}, 3},
		{"", args{[]int{3, 2, 2, 3}, 2}, 2},
		{"", args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2}, 5},
		{"", args{[]int{2, 1, 2}, 2}, 1},
		{"", args{[]int{2, 2, 2}, 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement5(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeElement2([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	}
}
