package dp

import "testing"

// 494. Target Sum

func findTargetSumWays(nums []int, S int) int {
	// memo := map[[2]int]int{}
	res := 0
	findTargetsSumWaysHelper(nums, &res, 0, 0, S)
	return res

}

func findTargetsSumWaysHelper(nums []int, res *int, i, sum, s int) {
	if i == len(nums) {
		if sum == s {
			*res++
		}
		return
	}
	findTargetsSumWaysHelper(nums, res, i+1, sum+nums[i], s)
	findTargetsSumWaysHelper(nums, res, i+1, sum-nums[i], s)
}

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums []int
		S    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 1, 1, 1, 1}, 3}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.S); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
