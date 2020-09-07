package string

import (
	"reflect"
	"strconv"
	"testing"
)

// 93. Restore IP Addresses

// * one pass
//RRuntime: 0 ms, faster than 100.00% of Go online submissions for Restore IP Addresses.
//Memory Usage: 2.1 MB, less than 96.26% of Go online submissions for Restore IP Addresses.
func restoreIpAddresses(s string) []string {
	res := []string{}
	if len(s) == 0 {
		return res
	}
	parseIPAddress(s, &res, [4]byte{}, 0)
	return res
}

func parseIPAddress(s string, res *[]string, ip [4]byte, seg int) {
	if seg == 3 {
		if len(s) > 1 && s[0] == '0' {
			return
		}
		t, _ := strconv.Atoi(s)
		if t <= 255 {
			ip[seg] = byte(t)
			*res = append(*res, byteToIPString(ip))
		}
		return
	}

	if len(s) > 1 {
		ip[seg] = s[0] - 48
		parseIPAddress(s[1:], res, ip, seg+1)

	}
	if len(s) > 2 && s[0] != '0' {
		t, _ := strconv.Atoi(s[:2])
		ip[seg] = byte(t)
		parseIPAddress(s[2:], res, ip, seg+1)
	}
	if len(s) > 3 && s[0] != '0' {
		t, _ := strconv.Atoi(s[:3])
		if t < 256 {
			ip[seg] = byte(t)
			parseIPAddress(s[3:], res, ip, seg+1)
		}
	}
}

func byteToIPString(b [4]byte) string {
	var res string
	for i, v := range b {
		if i != 0 {
			res += "."
		}
		res += strconv.Itoa(int(v))
	}
	return res
}

func Test_byteToIPString(t *testing.T) {
	a := [4]byte{123, 23, 4, 0}
	r := byteToIPString(a)
	t.Log(r)
}

func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"00000", args{"00000"}, []string{}},
		{"", args{""}, []string{}},
		{"101023", args{"101023"}, []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}},
		{"010010", args{"010010"}, []string{"0.10.0.10", "0.100.1.0"}},
		{"25525511135", args{"25525511135"}, []string{"255.255.11.135", "255.255.111.35"}},
		{"1234", args{"1234"}, []string{"1.2.3.4"}},
		{"0000", args{"0000"}, []string{"0.0.0.0"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_restoreIpAddresses(b *testing.B) {
	a := "25525511135"
	for i := 0; i < b.N; i++ {
		_ = restoreIpAddresses(a)
	}
}
