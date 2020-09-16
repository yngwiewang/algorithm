package array

import "testing"

// 26. Remove Duplicates from Sorted Array
// 双指针，笨
func removeDuplicatesFromSortedArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 1
	for i < len(nums) && nums[i] != nums[i-1] {
		i++
	}
	if i == len(nums) {
		return len(nums)
	}
	j := i + 1
	tmp := nums[i]
	for j < len(nums) {
		if nums[j] == tmp {
			j++
		} else {
			tmp = nums[j]
			nums[i], nums[j] = nums[j], nums[i]
			i, j = i+1, j+1
		}
	}
	return i
}

// 简单双指针
func removeDuplicatesFromSortedArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

// 参照最快答案
func removeDuplicatesFromSortedArray2(nums []int) int {
	length := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[length] {
			length++
			nums[length] = nums[i]
		}
	}
	return length + 1
}

func Test_removeDuplicatesFromSortedArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"[1,2]", args{[]int{1, 2}}, 2},
		{"[0,0,1,1,1,2,2,3,3,4]", args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, 5},
		{"[1,1,2]", args{[]int{1, 1, 2}}, 2},
		{"[1,1]", args{[]int{1, 1}}, 1},
		{"[1,2,2]", args{[]int{1, 2, 2}}, 2},
		{"[0,1,2,2,2,3,3,4,5,5,5,5,6,7,7,7,8,8,9,9,9,10]", args{[]int{0, 1, 2, 2, 2, 3, 3, 4, 5, 5, 5, 5, 6, 7, 7, 7, 8, 8, 9, 9, 9, 10}}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicatesFromSortedArray2(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
