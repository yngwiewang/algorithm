package dp

import (
	"reflect"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 474. Ones and Zeroes

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][][]int, len(strs)+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}
	for i := 1; i <= len(strs); i++ {
		cur := count(strs[i-1])
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				if cur[0] > j || cur[1] > k {
					dp[i][j][k] = dp[i-1][j][k]
				} else {
					dp[i][j][k] = common.MaxInt(dp[i-1][j][k],
						dp[i-1][j-cur[0]][k-cur[1]]+1)
				}
			}
		}
	}
	return dp[len(strs)][m][n]
}

func count(s string) (res [2]int) {
	for _, v := range s {
		res[v-48]++
	}
	return
}

func Test_count(t *testing.T) {
	type args struct {
		in0 string
	}
	tests := []struct {
		name    string
		args    args
		wantRes [2]int
	}{
		{"11000", args{"11000"}, [2]int{3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := count(tt.args.in0); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("count() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_findMaxForm(t *testing.T) {
	type args struct {
		strs []string
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"10 0001 111001 1 0", args{[]string{"10", "0001", "111001", "1", "0"}, 5, 3}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxForm(tt.args.strs, tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("findMaxForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
