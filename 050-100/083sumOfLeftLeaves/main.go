package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func test() *TreeNode {
	n5 := &TreeNode{
		Val:   5,
	}
	n4 := &TreeNode{
		Val:   4,
	}
	n3 := &TreeNode{
		Val:   3,
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

func test2() *TreeNode {
	n7 := &TreeNode{
		Val:   7,
	}
	n15 := &TreeNode{
		Val:   15,
	}
	n20 := &TreeNode{
		Val:   20,
		Left:n15,
		Right:n7,
	}
	n9 := &TreeNode{
		Val:   9,
	}
	n3 := &TreeNode{
		Val:   3,
		Left:n9,
		Right:n20,
	}
	return n3
}

func main()  {
	fmt.Println(sumOfLeftLeaves(test()))
}

// 404 左叶子之和 左叶子节点，非左节点
func sumOfLeftLeaves(root *TreeNode) int {
	value := 0
	if root == nil {
		return 0
	}
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			value += root.Left.Val
		}
		 value += sumOfLeftLeaves(root.Left)
	}
	if root.Right != nil {
		value += sumOfLeftLeaves(root.Right)
	}
	return value
}

