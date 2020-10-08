package array

import (
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 152. Maximum Product Subarray
// DP
func maxProduct(nums []int) int {
	res := nums[0]

	for i, imax, imin := 1, res, res; i < len(nums); i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		imax = common.MaxInt(nums[i], nums[i]*imax)
		imin = common.MinInt(nums[i], nums[i]*imin)
		res = common.MaxInt(res, imax)
	}
	return res
}

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2,3,-2,4", args{[]int{2, 3, -2, 4}}, 6},
		{"", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
