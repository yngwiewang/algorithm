package array

import "testing"

func subarraySum(nums []int, k int) int {
	var count, sum int
	seen := make(map[int]int)
	for _, n := range nums {
		sum += n
		if sum == k {
			count++
		}
		count += seen[sum-k]
		seen[sum]++
	}
	return count
}

func TestSubarraySum(t *testing.T) {
	a := []int{4, 3, 1, 6, 1, 6, 0, 7, 0, 0, 0}
	res := subarraySum1(a, 7)
	t.Log(res)
}

func subarraySum1(nums []int, k int) int {
	var count, sum int
	seen := make(map[int]int)
	seen[0] = 1
	for _, n := range nums {
		sum += n
		if tmp, ok := seen[sum-k]; ok {
			count += tmp
		}

		seen[sum]++
	}
	return count
}
