package array

import "testing"

// 713. Subarray Product Less Than K
// sliding window
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	prod := 1
	res, l := 0, 0
	for r, v := range nums {
		prod *= v
		for prod >= k {
			prod /= nums[l]
			l++
		}
		res += r - l + 1
	}
	return res
}

func Test_numSubarrayProductLessThanK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"10,5,2,6", args{[]int{10, 5, 2, 6}, 100}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayProductLessThanK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("numSubarrayProductLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 循环遍历 TLE
func numSubarrayProductLessThanKA(nums []int, k int) int {
	res := 0
	for i, v := range nums {
		product := v
		if product < k {
			res++
			j := i + 1
			for ; nums[j] == 1 && j < len(nums); res, j = res+1, j+1 {
			}
			for ; j < len(nums); j++ {
				product *= nums[j]
				if product < k {
					res++
				} else {
					break
				}
			}
		}
	}
	return res
}
