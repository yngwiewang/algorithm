package string

import (
	"strconv"
	"testing"
)

// 91. Decode Ways

// * 递归，如果记录编码的话非常低效:
// * Runtime: 4456 ms, faster than 5.18% of Go online submissions for Decode Ways.
// * Memory Usage: 6.8 MB, less than 5.18% of Go online submissions for Decode Ways.
func numDecodings(s string) int {
	encoding := map[byte]byte{}
	for i := 1; i <= 26; i++ {
		encoding[byte(i)] = byte(64 + i)
	}
	decodings := []byte{}
	var res int
	dfsDecoding(encoding, decodings, s, &res)
	return res
}

func dfsDecoding(e map[byte]byte, d []byte, s string, res *int) {
	if len(s) == 0 {
		*res++
		return
	} else if s[0] == byte('0') {
		return
	}

	d = append(d, e[s[0]-48])
	dfsDecoding(e, d, s[1:], res)
	d = d[:len(d)-1]

	if len(s) >= 2 {
		k, _ := strconv.Atoi(s[:2])
		if k >= 0 && k <= 26 {
			d = append(d, e[byte(k)])
			dfsDecoding(e, d, s[2:], res)
			if len(d) > 2 {
				d = d[:len(d)-2]
			}
		}
	}
}

/* 递归，只计数，快了不少
Runtime: 920 ms, faster than 13.55% of Go online submissions for Decode Ways.
Memory Usage: 2.2 MB, less than 41.43% of Go online submissions for Decode Ways.
*/
func numDecodings1(s string) int {
	var res int
	dfsDecoding1(s, &res)
	return res
}

func dfsDecoding1(s string, res *int) {
	if len(s) == 0 {
		*res++
		return
	} else if s[0] == byte('0') {
		return
	}

	dfsDecoding1(s[1:], res)

	if len(s) >= 2 {
		k, _ := strconv.Atoi(s[:2])
		if k >= 0 && k <= 26 {
			dfsDecoding1(s[2:], res)
		}
	}
}

// 动态规划，4ms
func numDecodings2(s string) int {
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s)+1)

	dp[0], dp[1] = 1, 1

	for i := 2; i < len(s)+1; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		twoDigit, _ := strconv.Atoi(s[i-2 : i])
		if twoDigit >= 10 && twoDigit <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

// ! 优化dp，初始判断非常tricky
func numDecodings3(s string) int {
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	n0, n1 := 1, 0
	if s[1] != '0' {
		n1 = 1
	}

	if s[:2] < "27" {
		n1 += n0
	}
	for i := 2; i < len(s); i++ {
		var t int
		if s[i] != '0' {
			t = n1
		}
		if s[i-1] != '0' && s[i-1:i+1] < "27" {
			t += n0
		}
		if t == 0 {
			return 0
		}
		n0, n1 = n1, t
	}

	return n1
}

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"3534532543253245", args{"3534532543253245"}, 8},
		{"30", args{"30"}, 0},
		{"301", args{"301"}, 0},
		{"101", args{"101"}, 1},
		{"100", args{"100"}, 0},
		{"10", args{"10"}, 1},
		{"1", args{"1"}, 1},
		{"226", args{"226"}, 3},
		{"17", args{"17"}, 2},
		{"01", args{"01"}, 0},
		{"12394329575982", args{"12394329575982"}, 3},
		{"12039403295075982", args{"12039403295075982"}, 0},
		{"12", args{"12"}, 2},
		{"1", args{"1"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings3(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
