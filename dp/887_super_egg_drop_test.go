package dp

import (
	"math"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 887. Super Egg Drop
// dp, recursive, TLE
func superEggDrop(K int, N int) int {
	memo := make([][]int, K+1)
	for i := range memo {
		memo[i] = make([]int, N+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return superEggDropHelper(K, N, memo)
}

func superEggDropHelper1(k, n int, memo [][]int) int {
	if k == 1 {
		return n
	}
	if n == 0 {
		return 0
	}
	if memo[k][n] != -1 {
		return memo[k][n]
	}
	res := math.MaxInt32
	for i := 1; i <= n; i++ {
		res = common.MinInt(res, 1+common.MaxInt(superEggDropHelper(k-1, i-1, memo),
			superEggDropHelper(k, n-i, memo)))
	}
	memo[k][n] = res
	return memo[k][n]
}

// binary search optimize
func superEggDropHelper(k, n int, memo [][]int) int {
	if k == 1 {
		return n
	}
	if n == 0 {
		return 0
	}
	if memo[k][n] != -1 {
		return memo[k][n]
	}
	res := math.MaxInt32
	l, h := 1, n
	for l <= h {
		m := (l + h) / 2
		broken := superEggDropHelper(k-1, m-1, memo)
		notBroken := superEggDropHelper(k, n-m, memo)
		if broken > notBroken {
			h = m - 1
			res = common.MinInt(res, broken+1)
		} else {
			l = m + 1
			res = common.MinInt(res, notBroken+1)
		}
	}

	memo[k][n] = res
	return memo[k][n]
}

func Test_superEggDrop(t *testing.T) {
	type args struct {
		K int
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1, 2", args{1, 2}, 2},
		{"2, 6", args{2, 6}, 3},
		{"3, 14", args{3, 14}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superEggDrop(tt.args.K, tt.args.N); got != tt.want {
				t.Errorf("superEggDrop() = %v, want %v", got, tt.want)
			}
		})
	}
}
