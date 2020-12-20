package dp

import (
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 213. House Robber II

func robCircle(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return common.MaxInt(rob(nums[:len(nums)-1]), rob(nums[1:]))
}

func Test_robCircle(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1,2", args{[]int{1, 2}}, 2},
		{"1", args{[]int{1}}, 1},
		{"2,3,2", args{[]int{2, 3, 2}}, 3},
		{"1,2,3,1", args{[]int{1, 2, 3, 1}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robCircle(tt.args.nums); got != tt.want {
				t.Errorf("robCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}
