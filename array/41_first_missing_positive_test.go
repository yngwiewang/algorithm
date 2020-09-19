package array

import (
	"sort"
	"testing"
)

// 41. First Missing Positive

/** 先排序，时间复杂度不满足O(n)的要求，但是仍然达到了100%
 * 边界情况很多，需要小心处理 */

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	i := 0
	for i < len(nums) && nums[i] <= 0 {
		i++
	}
	copy(nums, nums[i:])

	j := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] != j {
			return j
		}
		if i < len(nums)-1 && nums[i+1] == nums[i] {
			continue
		}
		j++
	}
	return j
}

// 用map，非O(1)空间复杂度
func firstMissingPositive1(ns []int) int {
	m := make(map[int]bool)
	for _, v := range ns {
		m[v] = true
	}
	
	for i:=1; i < len(m)+1; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return len(m)+1
}

// 正确答案，O(n)，比较奇妙
func firstMissingPositive2(ns []int) int {

	// to cater for len(nums) < 2
	ns = append(ns, 0)

	for i := 0; i < len(ns); i++ {
		if ns[i] < 0 || ns[i] >= len(ns) {
			ns[i] = 0
		}
	}

	for i := 0; i < len(ns); i++ {
		ns[ns[i]%len(ns)] += len(ns)
	}

	for i := 1; i < len(ns); i++ {
		if ns[i] < len(ns) {
			return i
		}
	}

	return len(ns)

}

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1,1,1,2,2,3,3,3,3,4,4,5", args{[]int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 4, 5}}, 6},
		{"[0,2,2,1,1]", args{[]int{0, 2, 2, 1, 1}}, 3},
		{"-10,-3,-100,-1000,-239,1", args{[]int{-10, -3, -100, -1000, -239, 1}}, 2},
		{"1,1,1,2,2,3,3,3,3,5", args{[]int{1, 1, 1, 2, 2, 3, 3, 3, 3, 5}}, 4},
		{"9,1", args{[]int{9, 1}}, 2},
		{"1,2,0", args{[]int{1, 2, 0}}, 3},
		{"9,2", args{[]int{9, 2}}, 1},
		{"0", args{[]int{0}}, 1},
		{"7,8,9,11,12", args{[]int{7, 8, 9, 11, 12}}, 1},
		{"3,4,-1,1", args{[]int{3, 4, -1, 1}}, 2},
		{"-1,-2", args{[]int{-1, -2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive1(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
