package array

import "testing"

// 53. Maximum Subarray


func maxSubArray(nums []int) int {
	max, curMax := -1<<31, -1<<31
	for _, v := range nums {
		if v+curMax > v {
			curMax = v + curMax
		} else {
			curMax = v
		}
		if curMax > max {
			max = curMax
		}
	}
	return max
}
func maxSubArray1(nums []int) int {
	max := nums[0]
	for i := 1;i<len(nums);i++ {		
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"-2,1,-3,4,-1,2,1,-5,4", args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, 6},
		{"1", args{[]int{1}}, 1},
		{"0", args{[]int{0}}, 0},
		{"-1", args{[]int{-1}}, -1},
		{"-2147483647", args{[]int{-2147483647}}, -2147483647},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray1(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
