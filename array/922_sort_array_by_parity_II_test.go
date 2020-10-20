package array

import "testing"

// 922. Sort Array By Parity II
func sortArrayByParityIIA(A []int) []int {
	for i, j := 0, 1; i < len(A) && j < len(A); {
		if A[i]%2 == 0 {
			i += 2
		} else if A[j]%2 == 1 {
			j += 2
		} else {
			A[i], A[j] = A[j], A[i]
			i, j = i+2, j+2
		}
	}
	return A
}

func sortArrayByParityII(A []int) []int {
	for i, j := 0, 1; i < len(A); i += 2 {
		if A[i]%2 != 0 {
			for ; A[j]%2 != 0; j += 2 {
			}

			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}

func Test_sortArrayParityII(t *testing.T) {
	a := []int{4, 2, 5, 7}
	b := sortArrayByParityII(a)
	t.Log(b)
}

func Benchmark_sortArrayParityII(b *testing.B) {
	a := []int{4, 2, 5, 7}
	for i := 0; i < b.N; i++ {
		sortArrayByParityIIA(a)
	}
}
