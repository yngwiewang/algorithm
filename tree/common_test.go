package tree

import (
	"fmt"
	"testing"
)

func TestArrayToBinaryTree(t *testing.T) {
	arrays := [][]int{
		{},
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}
	for _, a := range arrays {
		t := arrayToBinaryTree(a)
		fmt.Println(preorderTraversal(t))
	}
}
