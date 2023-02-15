package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func test() *TreeNode {
	n6 := &TreeNode{
		Val:   6,
	}
	n5 := &TreeNode{
		Val:   5,
	}
	n4 := &TreeNode{
		Val:   4,
	}
	n3 := &TreeNode{
		Val:   3,
		Left:n6,
	}
	n2 := &TreeNode{
		Val:   2,
		Left:n4,
		Right:n5,
	}
	n1 := &TreeNode{
		Val:   1,
		Left:n2,
		Right:n3,
	}
	return n1
}

func main()  {
	fmt.Println(minDepth(test()))
}

// 111 二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right != nil {
		return 1 + int(math.Min(float64(minDepth(root.Left)), float64(minDepth(root.Right))))
	}else if root.Left == nil {   //左空，右空，左右空
		return 1 + minDepth(root.Right)
	}else {
		return 1 + minDepth(root.Left)
	}
}

