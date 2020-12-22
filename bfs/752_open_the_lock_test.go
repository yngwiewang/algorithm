package bfs

import (
	"testing"
)

// 752. Open the Lock
// BFS
func openLock(deadends []string, target string) int {
	q := []string{"0000"}
	memo := make(map[string]bool, len(deadends))
	for _, v := range deadends {
		if v == "0000" {
			return -1
		}
		memo[v] = true
	}
	res := 0
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			s := q[i]
			if s == target {
				return res
			}
			for j := 0; j <= 3; j++ {
				turned := turnUp(s, j)
				if _, ok := memo[turned]; !ok {
					memo[turned] = true
					q = append(q, turned)
				}
				turned = turnDown(s, j)
				if _, ok := memo[turned]; !ok {
					memo[turned] = true
					q = append(q, turned)
				}

			}
		}
		q = q[size:]
		res++
	}
	return -1
}

func turnUp(s string, i int) string {
	b := []byte(s)
	if b[i] == '9' {
		b[i] = '0'
	} else {
		b[i]++
	}
	return string(b)
}

func turnDown(s string, i int) string {
	b := []byte(s)
	if b[i] == '0' {
		b[i] = '9'
	} else {
		b[i]--
	}
	return string(b)
}

func Test_turnUp(t *testing.T) {
	a := "0900"
	t.Log(turnUp(a, 1))
}

func Test_turnDown(t *testing.T) {
	a := "0900"
	t.Log(turnDown(a, 2))
}

func Test_openLock(t *testing.T) {
	type args struct {
		deadends []string
		target   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"0201", "0101", "0102", "1212", "2002"}, "0202"}, 6},
		{"2", args{[]string{"8888"}, "0009"}, 1},
		{"3", args{[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"}, -1},
		{"4", args{[]string{"0000"}, "8888"}, -1},
		{"5", args{[]string{"1000", "0100", "0010", "0001", "9000", "0900", "0090", "0009"}, "0202"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := openLock(tt.args.deadends, tt.args.target); got != tt.want {
				t.Errorf("openLock() = %v, want %v", got, tt.want)
			}
		})
	}
}
