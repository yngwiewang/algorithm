package stack

import (
	"strings"
	"testing"
)

// 394. Decode String
func decodeString(s string) string {
	var (
		n   int
		res string
		str []string
		num []int
	)
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= '0' && s[i] <= '9':
			n = n*10 + int(s[i]-'0')
		case s[i] == '[':
			num = append(num, n)
			n = 0
			str = append(str, res)
			res = ""
		case s[i] == ']':
			tmpS := str[len(str)-1]
			str = str[:len(str)-1]
			tmpN := num[len(num)-1]
			num = num[:len(num)-1]
			res = tmpS + strings.Repeat(res, tmpN)

		default:
			res += string(s[i])
		}
		// fmt.Println(num, str, n, res)
	}
	return res
}

func Test_decodeString(t *testing.T) {
	x := "3[a]2[bc]"
	s := decodeString(x)
	t.Log(s)
}

func Test_str(t *testing.T) {
	a, b := 'a', 'b'
	t.Log(string(a + b))
}
