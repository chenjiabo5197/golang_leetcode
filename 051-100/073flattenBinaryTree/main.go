package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	n6 := &TreeNode{
		Val:   6,
	}
	n4 := &TreeNode{
		Val:   4,
	}
	n3 := &TreeNode{
		Val:   3,
	}
	n5 := &TreeNode{
		Val:   5,
		Right:n6,
	}
	n2 := &TreeNode{
		Val:   2,
		Left:n3,
		Right:n4,
	}
	n1 := &TreeNode{
		Val:   1,
		Left:n2,
		Right:n5,
	}
	flatten(n1)
	fmt.Println(n1)
}

// 114 二叉树展开为链表
func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	if root.Left != nil {
		flatten(root.Left)
	}
	if root.Right != nil {
		flatten(root.Right)
	}
	//if root.Left != nil {
	//	root.Left.Right = root.Right
	//}
	//if root.Left == nil {
	//	root.Left = root.Right
	//}
	//root.Right = root.Left
	//root.Left = nil
	if root.Left == nil {
		return
	}
	tempNode := &TreeNode{
		Left: root.Left,
	}
	for tempNode.Left != nil {
		tempNode = tempNode.Left
	}
	for tempNode.Right != nil {
		tempNode = tempNode.Right
	}
	tempNode.Right = root.Right
	root.Right = root.Left
	root.Left = nil
	return
}
