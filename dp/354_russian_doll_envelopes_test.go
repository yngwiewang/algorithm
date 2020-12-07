package dp

import (
	"fmt"
	"sort"
	"testing"

	"github.com/yngwiewang/algorithm/common"
)

// 300. Longest Increasing Subsequence
// dp table
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	sort.Slice(envelopes, func(i, j int) bool {
		// if envelopes[i][0] < envelopes[j][0] {
		// 	return envelopes[i][0] < envelopes[j][0]
		// }
		// return envelopes[i][1] < envelopes[j][1]
		return envelopes[i][0] < envelopes[j][0]
	})
	fmt.Println(envelopes)
	dp := make([]int, len(envelopes))
	res := 1
	for i := range envelopes {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[j][0] < envelopes[i][0] &&
				envelopes[j][1] < envelopes[i][1] {
				dp[i] = common.MaxInt(dp[i], 1+dp[j])
				res = common.MaxInt(res, dp[i])
			}
		}
	}
	fmt.Println(dp)
	return res
}

func Test_maxEnvelopes(t *testing.T) {
	type args struct {
		envelopes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"[[10,4],[13,18],[1,5],[13,15],[3,12],[12,11],[17,15],[7,1],[17,18],[7,19],[2,5],[8,9],[18,10],[7,6],[17,7]]", args{[][]int{{10, 4}, {13, 18}, {1, 5}, {13, 15}, {3, 12}, {12, 11}, {17, 15}, {7, 6}, {17, 18}, {7, 19}, {2, 5}, {8, 9}, {18, 10}, {7, 1}, {17, 7}}}, 6},
		{"[[10,17],[10,19],[16,2],[19,18],[5,6]]]", args{[][]int{{10, 17}, {10, 19}, {16, 2}, {19, 18}, {5, 6}}}, 3},
		{"[[5,4],[6,4],[6,7],[2,3]]", args{[][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}}, 3},
		{"[[4,6],[3,7],[4,8],[2,3]]", args{[][]int{{4, 6}, {3, 7}, {4, 8}, {2, 3}}}, 3},
		{"[[]]", args{[][]int{{}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}
}
