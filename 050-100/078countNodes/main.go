package main

import "fmt"

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
	fmt.Println(countNodes(test()))
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
