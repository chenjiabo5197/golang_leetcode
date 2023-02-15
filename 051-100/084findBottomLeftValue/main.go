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
	n7 := &TreeNode{
		Val:   7,
	}
	n6 := &TreeNode{
		Val:   6,
	}
	n5 := &TreeNode{
		Val:   5,
		Left:n7,
	}
	n4 := &TreeNode{
		Val:   4,
	}
	n3 := &TreeNode{
		Val:   3,
		Left:n5,
		Right:n6,
	}
	n2 := &TreeNode{
		Val:   2,
		Left:n4,
	}
	n1 := &TreeNode{
		Val:   1,
		Left:n2,
		Right:n3,
	}
	return n1
}

func main()  {
	fmt.Println(findBottomLeftValue(test()))
}

// 513 找树左下角的值
func findBottomLeftValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	if getDepth(root.Left) >= getDepth(root.Right) {
		return findBottomLeftValue(root.Left)
	}else {
		return findBottomLeftValue(root.Right)
	}
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + int(math.Max(float64(getDepth(root.Left)), float64(getDepth(root.Right))))
}
