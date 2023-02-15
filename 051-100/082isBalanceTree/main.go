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
	n41 := &TreeNode{
		Val:   4,
	}
	n42 := &TreeNode{
		Val:   4,
	}
	n31 := &TreeNode{
		Val:   3,
		Left:n41,
		Right:n42,
	}
	n32 := &TreeNode{
		Val:   3,
	}
	n21 := &TreeNode{
		Val:   2,
		Left:n31,
		Right:n32,
	}
	n22 := &TreeNode{
		Val:   2,
	}
	n1 := &TreeNode{
		Val:   1,
		Left:n21,
		Right:n22,
	}
	return n1
}

func main()  {
	fmt.Println(isBalanced(test()))
}

// 110平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if math.Abs(float64(leftDepth) - float64(rightDepth)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
}

