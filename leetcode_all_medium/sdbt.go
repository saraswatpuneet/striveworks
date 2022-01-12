package main

import (
	"strconv"
	"strings"
)

// serialize and deserialize binary tree
// serialize: serialize a binary tree to a string
// deserialize: deserialize a string to a binary tree
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	queue := []*TreeNode{root}
	cstring := []string{strconv.Itoa(root.Val)}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
				cstring = append(cstring, strconv.Itoa(node.Left.Val))
			} else {
				cstring = append(cstring, "null")
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				cstring = append(cstring, strconv.Itoa(node.Right.Val))
			} else {
				cstring = append(cstring, "null")
			}
		}
	}
	return strings.Join(cstring, ",")
}

func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	cstring := strings.Split(data, ",")
	queue := []*TreeNode{}
	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(cstring[0])
	queue = append(queue, root)
	counter := 1
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if cstring[counter] != "null" {
				left := &TreeNode{}
				left.Val, _ = strconv.Atoi(cstring[counter])
				node.Left = left
				queue = append(queue, left)
			}
			counter++
			if cstring[counter] != "null" {
				right := &TreeNode{}
				right.Val, _ = strconv.Atoi(cstring[counter])
				node.Right = right
				queue = append(queue, right)
			}
			counter++
		}
	}
	return root
}
