package dp

import "testing"

// 509. Fibonacci Number
// recursive with memo
func fib(N int) int {
	if N <= 1 {
		return N
	}
	memo := make([]int, N+1)
	memo[1], memo[2] = 1, 1
	res := fibHelper(N, memo)
	return res
}

func fibHelper(n int, memo []int) int {
	if memo[n] != 0 {
		return memo[n]
	}
	return fibHelper(n-1, memo) + fibHelper(n-2, memo)
}

// dp array
func fib1(N int) int {
	if N <= 1 {
		return N
	}
	dp := make([]int, N+1)
	dp[1], dp[2] = 1, 1
	for i := 3; i <= N; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[N]
}

// dp int
func fib2(N int) int {
	if N <= 1 {
		return N
	}
	a, b := 1, 1
	for i := 0; i < N-2; i++ {
		b, a = a+b, b
	}
	return b
}

func Test_fib(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{0}, 0},
		{"1", args{1}, 1},
		{"2", args{1}, 1},
		{"3", args{3}, 2},
		{"4", args{4}, 3},
		{"6", args{6}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib2(tt.args.N); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
