package array

import "testing"

// 217. Contains Duplicate
// use map
func containsDuplicateA(n []int) bool {
	m := make(map[int]struct{})

	for _, v := range n {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = struct{}{}
	}
	return false
}

// 牺牲空间换取更快的效率
func containsDuplicate(n []int) bool {
	m := make(map[int]struct{}, len(n))

	for _, v := range n {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = struct{}{}
	}
	return false
}

func Benchmark_containsDuplicate(b *testing.B) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		containsDuplicate(a)
	}
}

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1,2,3", args{[]int{1, 2, 3}}, false},
		{"1,2,1", args{[]int{1, 2, 1}}, true},
		{"1", args{[]int{1}}, false},
		{"", args{[]int{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.n); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
