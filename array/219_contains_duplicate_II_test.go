package array

import "testing"

// 219. Contains Duplicate II
// 遍历，龟速
func containsNearbyDuplicateA(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}
	i, j := 0, 1
	for {
		if i == len(nums)-1 {
			return false
		}
		for j < len(nums) && nums[i] != nums[j] {
			j++
		}
		if j == len(nums) {
			i, j = i+1, i+2
			continue
		}
		if j-i <= k {
			return true
		} else {
			i, j = i+1, i+2
		}
	}
}

// use map, more faster
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}
	m := make(map[int]int, len(nums))
	for i, v := range nums{
		if _, ok := m[v]; ok && i - m[v]<=k{
			return true
		}
		m[v] = i
	}
	return false
}

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1,2,3,1,2,3", args{[]int{1, 2, 3, 1, 2, 3}, 2}, false},
		{"1,2,3,1", args{[]int{1, 2, 3, 1}, 3}, true},
		{"1,0,1,1", args{[]int{1, 0, 1, 1}, 1}, true},
		{"4,1,2,3,1,5", args{[]int{4, 1, 2, 3, 1, 5}, 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
