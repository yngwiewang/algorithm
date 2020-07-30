package array

import (
	"fmt"
	"sort"
	"testing"
)

// use set, TLE
func threeSum(nums []int) [][]int {
	var res [][]int
	m := make(map[int][]int)
	for i, e := range nums {
		m[e] = append(m[e], i)
	}
	seen := make(map[[3]int]bool)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if k, ok := m[0-nums[i]-nums[j]]; ok {
				for _, p := range k {
					if p > i && p > j {
						if notDup(seen, nums[i], nums[j], 0-nums[i]-nums[j]) {
							res = append(res, []int{nums[i], nums[j], 0 - nums[i] - nums[j]})
							seen[[3]int{nums[i], nums[j], 0 - nums[i] - nums[j]}] = true
						}
					}
				}
			}
		}

	}
	return res
}

func notDup(m map[[3]int]bool, a, b, c int) bool {
	if m[[3]int{a, b, c}] == true || m[[3]int{a, c, b}] == true ||
		m[[3]int{b, a, c}] == true || m[[3]int{b, c, a}] == true ||
		m[[3]int{c, a, b}] == true || m[[3]int{c, b, a}] == true {
		return false
	}
	return true
}

func TestThreeSum(t *testing.T) {
	ls := [][]int{{-1, -2, -3, 4, 1, 3, 0, 3, -2, 1, -2, 2, -1, 1, -5, 4, -3}, {-1, 0, 1, 2, -1, -4}, {1, -1, -1, 0}}
	for _, l := range ls {
		res := threeSum1(l)
		fmt.Println(res)
	}
}

func threeSum1(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	i := 0
	for i < len(nums) && nums[i] <= 0 {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for r > 0 && nums[r] == nums[r-1] {
					r--
				}
				r--
			} else if sum > 0 {
				for r > 0 && nums[r] == nums[r-1] {
					r--
				}
				r--
			} else if sum < 0 {
				for l < len(nums)-1 && nums[l] == nums[l+1] {
					l++
				}
				l++
			}
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] && nums[i] <= 0 {
			i++
		}
		i++
	}
	return res
}
