package array

import "testing"

// 907. Sum of Subarray Minimums
// brute
func sumSubarrayMins(a []int) int {
	res := 0
	for i := 0; i < len(a)-1; i++ {
		min := a[i]
		res += min
		for j := i + 1; j < len(a); j++ {
			if a[j] < min {
				min = a[j]
			}
			res += min
		}
	}
	res += a[len(a)-1]
	return res % (1000000007)
}

func Test_sumSubarrayMins(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"3,1,2,4", args{[]int{3, 1, 2, 4}}, 17},
		{"2,1,3", args{[]int{2,1,3}}, 9},
		{"1", args{[]int{1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSubarrayMins(tt.args.a); got != tt.want {
				t.Errorf("sumSubarrayMins() = %v, want %v", got, tt.want)
			}
		})
	}
}
