package array

import (
	"math"
	"testing"
)

// 121. Best Time to Buy and Sell Stock
func maxProfit1(prices []int) int {	
	if len(prices) == 0{
		return 0
	}

	min := prices[0]
	profit := 0
	
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > profit {
			profit = prices[i] - min
		}
	}
	return profit
}

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
		{"1", args{[]int{7, 6, 4, 3, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 复习
func maxProfit(prices []int) int {	
	minPrice, maxProfit := math.MaxInt32, 0
	for i:=0;i<len(prices);i++{
		if prices[i] < minPrice{
			minPrice = prices[i]
		}
		if prices[i] - minPrice > maxProfit{
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}
