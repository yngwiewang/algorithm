package string

import (
	"container/heap"
	"fmt"
	"testing"
)

// 767. Reorganize String
// 参考最好的答案做的，很巧妙，不用排序，时空复杂度都战胜100%
func reorganizeString(S string) string {
	var cc [26]int
	for _, v := range S {
		cc[v-'a']++
	}
	pos, max := maxChar(&cc)
	if max > (len(S)+1)/2 {
		return ""
	}
	res := make([]byte, len(S))
	idx := 0
	for cc[pos] > 0 {
		res[idx] = byte(pos) + 'a'
		idx += 2
		cc[pos]--
	}
	for i := 0; i < 26; i++ {
		for cc[i] > 0 {
			if idx >= len(res) {
				idx = 1
			}
			res[idx] = byte(i) + 'a'
			idx += 2
			cc[i]--
		}
	}
	return string(res)
}

func maxChar(cc *[26]int) (pos, max int) {
	for i, v := range cc {
		if v > max {
			pos, max = i, v
		}
	}
	return
}

type pair struct {
	char  byte
	count int
}

type pairs []pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Less(i, j int) bool {
	return p[i].count > p[j].count
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *pairs) Push(x interface{}) {
	*p = append(*p, x.(pair))
}

func (p *pairs) Pop() (x interface{}) {
	x = (*p)[0]
	*p = (*p)[1:]
	return
}

func Test_heap(t *testing.T) {
	a := &pairs{pair{'c', 3}, pair{'a', 4}, pair{'b', 5}}
	fmt.Println(a)
	heap.Init(a)
	fmt.Println(a)
	x := a.Pop().(pair)
	fmt.Println(x)
	x.count -= 4
	a.Push(x)
	fmt.Println(a.Pop().(pair))
	fmt.Println(a.Pop().(pair))
	fmt.Println(a.Pop().(pair))
}

// 用堆实现
// ! 注意通过Less方法的大于号来设置大顶堆
// ! 入堆要用 heap.Push(h, elem) 而不是 h.Push(elem)
func reorganizeString1(S string) string {
	var count [26]int
	for i := range S {
		count[S[i]-'a']++
	}
	p := pairs{}
	for i, v := range count {
		if v != 0 {
			p = append(p, pair{byte(i) + 'a', v})
		}
	}
	heap.Init(&p)
	res := []byte{}
	for p.Len() > 1 {
		t1 := p.Pop().(pair)
		t2 := p.Pop().(pair)
		if len(res) > 0 && t1.char == res[len(res)-1] {
			t1, t2 = t2, t1
		}
		res = append(res, t1.char)
		res = append(res, t2.char)
		t1.count--
		t2.count--
		if t1.count > 0 {
			heap.Push(&p, t1)
		}
		if t2.count > 0 {
			heap.Push(&p, t2)
		}
		if p.Len() == 0 {
			break
		}
		if p.Len() == 1 {
			if (p)[0].count > 1 {
				return ""
			}
			res = append(res, (p)[0].char)
			break
		}
	}
	return string(res)
}

func Test_reorganizeString(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"gpneqthatplqrofqgwwfmhzxjddhyupnluzkkysofgqawjyrwhfgdpkhiqgkpupgdeonipvptkfqluytogoljiaexrnxckeofqojltdjuujcnjdjohqbrzzzznymyrbbcjjmacdqyhpwtcmmlpjbqictcvjgswqyqcjcribfmyajsodsqicwallszoqkxjsoskxxstdeavavnqnrjelsxxlermaxmlgqaaeuvneovumneazaegtlztlxhihpqbajjwjujyorhldxxbdocklrklgvnoubegjrfrscigsemporrjkiyncugkksedfpuiqzbmwdaagqlxivxawccavcrtelscbewrqaxvhknxpyzdzjuhvoizxkcxuxllbkyyygtqdngpffvdvtivnbnlsurzroxyxcevsojbhjhujqxenhlvlgzcsibcxwomfpyevumljanfpjpyhsqxxnaewknpnuhpeffdvtyjqvvyzjeoctivqwann", args{"gpneqthatplqrofqgwwfmhzxjddhyupnluzkkysofgqawjyrwhfgdpkhiqgkpupgdeonipvptkfqluytogoljiaexrnxckeofqojltdjuujcnjdjohqbrzzzznymyrbbcjjmacdqyhpwtcmmlpjbqictcvjgswqyqcjcribfmyajsodsqicwallszoqkxjsoskxxstdeavavnqnrjelsxxlermaxmlgqaaeuvneovumneazaegtlztlxhihpqbajjwjujyorhldxxbdocklrklgvnoubegjrfrscigsemporrjkiyncugkksedfpuiqzbmwdaagqlxivxawccavcrtelscbewrqaxvhknxpyzdzjuhvoizxkcxuxllbkyyygtqdngpffvdvtivnbnlsurzroxyxcevsojbhjhujqxenhlvlgzcsibcxwomfpyevumljanfpjpyhsqxxnaewknpnuhpeffdvtyjqvvyzjeoctivqwann"}, "jqjajljqjnxljqxnjexqjnxvlqajlceqzanxvpraolcecrqnxeycgjlvucrqnxeptaolvuszqkxeytgjdhucrqnxepmaolviszqkxfytgjdhucrbnxepmaolviszqkwfytgjdhucrbnxepmaolviszqkwfytgjdhucrbnxepmaolviszqkwfytgjdhucrbnxepmaolviszqkwfytgjdhucrbnxepmaolviszqkwfytgjdhucrbnxepmaolviszqkwfytgjdhucrbnxepmaolviszqkwfytgjdhucrbnxepmaolviszqkwfytgjdhucrbnxepmaolviszqkwfytgjdhucrbnxepmaolviszqkwfytgjdhucrbnxepmaolviszqkwfytgjdhucrbnxepmaolvisaqkwfyzgjdhucrknxeptovolsoibafyzqjnjhjcywyrpgudxsoikaypgvnlhjcbwpevudxjokayfgqnlhscpevdjol"},
		{"vvvlo", args{"vvvlo"}, "vovlv"},
		{"aabbcc", args{"aabbcc"}, "abcabc"},
		{"aabbccdd", args{"aabbccdd"}, "abcdabcd"},
		{"aab", args{"aab"}, "aba"},
		{"abbb", args{"abbb"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorganizeString1(tt.args.S); got != tt.want {
				t.Errorf("reorganizeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
