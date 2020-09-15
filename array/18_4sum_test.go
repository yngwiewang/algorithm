package array

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// 18. 4Sum

// 双指针，移动时跳过重复值，速度比较慢
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-3; {
		for j := i + 1; j < len(nums)-2; {
			a, b := j+1, len(nums)-1
			for a < b {
				sum := nums[a] + nums[b] + nums[i] + nums[j]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[a], nums[b]})
				}
				if sum < target {
					for a < len(nums)-1 {
						if nums[a] != nums[a+1] {
							a++
							break
						} else {
							a++
						}

					}
				} else {
					for b > 0 {
						if nums[b] != nums[b-1] {
							b--
							break
						} else {
							b--
						}
					}
				}
			}
			for j < len(nums)-2 {
				if nums[j] != nums[j+1] {
					j++
					break
				} else {
					j++
				}
			}
		}
		for i < len(nums)-3 {
			if nums[i] != nums[i+1] {
				i++
				break
			} else {
				i++
			}
		}
	}
	fmt.Println(res)
	return res
}

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"-4,0,2", args{[]int{-4, 0, 2}, 0}, [][]int{{-1, 0, 0, 1}}},
		{"-4,0,2,2", args{[]int{-4, 0, 2, 2}, 0}, [][]int{{-1, 0, 0, 1}}},
		{"-4,0,1,2,2,3", args{[]int{-4, 0, 2, 1, 2, 3}, 0}, [][]int{{-1, 0, 0, 1}}},
		{"1, 0, -1, 0, -2, 2", args{[]int{1, 0, -1, 0, -2, 2}, 0}, [][]int{{-1, 0, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
