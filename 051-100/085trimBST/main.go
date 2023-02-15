package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func test() *TreeNode {
	n1 := &TreeNode{
		Val:   1,
	}
	n2 := &TreeNode{
		Val:   2,
		Left:n1,
	}
	n0 := &TreeNode{
		Val:   0,
		Right:n2,
	}
	n4 := &TreeNode{
		Val:   4,
	}
	n3 := &TreeNode{
		Val:   3,
		Left:n0,
		Right:n4,
	}
	return n3
}

func main()  {
	low := 1
	high := 3
	fmt.Println(trimBST(test(), low, high))
}

// 669 修剪二叉搜索树
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	if root.Val < low {
		return trimBST(root.Right , low, high)
	}
	root.Right = trimBST(root.Right, low, high)
	root.Left = trimBST(root.Left, low, high)
	return root
}
