package array

import (
	"fmt"
	"testing"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxArea(height []int) int {
	a, b := 0, len(height)-1
	max := Min(height[a], height[b]) * (b - a)
	for a != b {
		if height[a] <= height[b] {
			a++
		} else {
			b--
		}
		max = Max(max, Min(height[a], height[b])*(b-a))
	}
	return max
}

func TestContainerWithMostWater(t *testing.T) {
	x := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	a := maxArea(x)
	fmt.Println(a)
}
