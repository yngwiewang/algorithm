package tree

import (
	"math"
	"testing"
)

// 230. Kth Smallest Element in a BST

// inorder traversal, use []int
func kthSmallest1(root *TreeNode, k int) int {
	inorder := []int{}
	kthSmallestHelper1(root, k, &inorder)
	return inorder[k-1]
}

func kthSmallestHelper1(root *TreeNode, k int, inorder *[]int) {
	if root == nil {
		return
	}
	kthSmallestHelper1(root.Left, k, inorder)
	*inorder = append(*inorder, root.Val)
	if len(*inorder) >= k {
		return
	}
	kthSmallestHelper1(root.Right, k, inorder)
}

// inorder traversal, use index
func kthSmallest(root *TreeNode, k int) int {
	res := math.MaxInt32
	i := 0
	kthSmallestHelper(root, k, &i, &res)
	return res
}

func kthSmallestHelper(root *TreeNode, k int, i, res *int) {
	if root == nil {
		return
	}
	kthSmallestHelper(root.Left, k, i, res)
	if *res != math.MaxInt32 {
		return
	}
	*i++
	if *i == k {
		*res = root.Val
	}
	kthSmallestHelper(root.Right, k, i, res)
}

// use stack
func kthSmallestStack(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

func Test_kthSmallest(t *testing.T) {
	c := Constructor()
	tree := c.deserialize("3,1,4,#,2,#,#,#,#,")
	res := kthSmallestStack(tree, 4)
	t.Log(res)
}

func Benchmark_kthSmallest1(b *testing.B) {
	c := Constructor()
	tree := c.deserialize("3,1,4,#,2,#,#,#,#,")
	for i := 0; i < b.N; i++ {
		kthSmallest(tree, 4)
	}
}
