package tree

import (
	"strconv"
	"strings"
	"testing"
)

// 297. Serialize and Deserialize Binary Tree

// preorder traversal
type Codec1 struct {
	sep   string
	null  string
	sb    strings.Builder
	nodes []*TreeNode
}

func Constructor1() Codec1 {
	return Codec1{
		sep:   ",",
		null:  "#",
		sb:    strings.Builder{},
		nodes: []*TreeNode{},
	}
}

func (codec1 *Codec1) preTravSer(root *TreeNode) {
	if root == nil {
		codec1.sb.WriteString(codec1.null)
		codec1.sb.WriteString(codec1.sep)
		return
	}
	codec1.sb.WriteString(strconv.Itoa(root.Val))
	codec1.sb.WriteString(codec1.sep)
	codec1.preTravSer(root.Left)
	codec1.preTravSer(root.Right)
}
func (codec1 *Codec1) preTravDeser() *TreeNode {
	root := codec1.nodes[0]
	codec1.nodes = codec1.nodes[1:]
	if root == nil {
		return nil
	}
	root.Left = codec1.preTravDeser()
	root.Right = codec1.preTravDeser()

	return root
}

// Serializes a tree to a single string.
func (codec1 *Codec1) serialize(root *TreeNode) string {
	codec1.preTravSer(root)
	return codec1.sb.String()
}

// Deserializes your encoded data to tree.
func (codec1 *Codec1) deserialize(data string) *TreeNode {
	ns := strings.Split(data, codec1.sep)
	codec1.nodes = make([]*TreeNode, len(ns)-1)
	for i := 0; i < len(ns)-1; i++ {
		if ns[i] == "#" {
			codec1.nodes[i] = nil
		} else {
			num, _ := strconv.Atoi(ns[i])
			codec1.nodes[i] = &TreeNode{num, nil, nil}
		}
	}

	return codec1.preTravDeser()
}

func Test_codec1(t *testing.T) {
	tree := arrayToBinaryTree([]int{1, 2, 3, 4, 5, 6})
	c := Constructor1()
	ser := c.serialize(tree)
	t.Log(ser)
	newTree := c.deserialize(ser)
	trav := preorderTraversal(newTree)
	t.Log(trav)
}

// level traversal, poor time complexity, better space complexity.
type Codec struct {
	sep   string
	null  string
	sb    strings.Builder
	nodes []*TreeNode
}

func Constructor() Codec {
	return Codec{
		sep:   ",",
		null:  "#",
		sb:    strings.Builder{},
		nodes: []*TreeNode{},
	}
}

// Serializes a tree to a single string.
func (codec *Codec) serialize(root *TreeNode) string {
	if root == nil {
		codec.sb.WriteString(codec.null)
		return codec.sb.String()
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		levelLen := len(q)
		for i := 0; i < levelLen; i++ {
			top := q[0]
			q = q[1:]
			if top == nil {
				codec.sb.WriteString(codec.null)
				codec.sb.WriteString(codec.sep)
				continue
			}
			codec.sb.WriteString(strconv.Itoa(top.Val))
			codec.sb.WriteString(codec.sep)

			q = append(q, top.Left)
			q = append(q, top.Right)
		}
	}
	return codec.sb.String()
}

// Deserializes your encoded data to tree.
func (codec *Codec) deserialize(data string) *TreeNode {
	if data == codec.null {
		return nil
	}
	ns := strings.Split(data, codec.sep)
	codec.nodes = make([]*TreeNode, len(ns)-1)
	for i := 0; i < len(ns)-1; i++ {
		if ns[i] == "#" {
			codec.nodes[i] = nil
		} else {
			num, _ := strconv.Atoi(ns[i])
			codec.nodes[i] = &TreeNode{num, nil, nil}
		}
	}
	q := []*TreeNode{codec.nodes[0]}
	nodeIndex := 0
	for len(q) > 0 {
		levelLen := len(q)
		for i := 0; i < levelLen; i++ {
			top := q[0]
			q = q[1:]
			nodeIndex++
			top.Left = codec.nodes[nodeIndex]
			if codec.nodes[nodeIndex] != nil {
				q = append(q, codec.nodes[nodeIndex])
			}
			nodeIndex++
			top.Right = codec.nodes[nodeIndex]
			if codec.nodes[nodeIndex] != nil {
				q = append(q, codec.nodes[nodeIndex])
			}
		}
	}
	return codec.nodes[0]
}

func Test_levelCodec(t *testing.T) {
	c := Constructor()
	ser := c.serialize(arrayToBinaryTree([]int{1, 2, 3, 4, 5}))
	t.Log(ser)
	deser := c.deserialize("1,2,3,#,#,4,5,6,7,#,#,#,#,#,#,")
	t.Log(preorderTraversal(deser))
	t.Log(preorderTraversal(c.deserialize("#")))
}
