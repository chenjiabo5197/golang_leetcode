package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	n3 := &TreeNode{
		Val:   3,
	}
	n2 := &TreeNode{
		Val:   2,
		Left:  n3,
	}
	n1 := &TreeNode{
		Val:   1,
		Right: n2,
	}
	fmt.Println("前序：", preorderTraversal(n1))
	fmt.Println("中序：", inorderTraversal(n1))
	fmt.Println("后序：", postorderTraversal(n1))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 前序遍历
func preorderTraversal(root *TreeNode) []int {
	var temp []int
	if root == nil {
		return temp
	}
	temp = append(temp, root.Val)
	temp = append(temp, preorderTraversal(root.Left)...)
	temp = append(temp, preorderTraversal(root.Right)...)
	return temp
}

//中序遍历
func inorderTraversal(root *TreeNode) []int {
	var temp []int
	if root == nil {
		return temp
	}
	temp = append(temp, inorderTraversal(root.Left)...)
	temp = append(temp, root.Val)
	temp = append(temp, inorderTraversal(root.Right)...)
	return temp
}

// 后序遍历
func postorderTraversal(root *TreeNode) []int {
	var temp []int
	if root == nil {
		return temp
	}
	temp = append(temp, postorderTraversal(root.Left)...)
	temp = append(temp, postorderTraversal(root.Right)...)
	temp = append(temp, root.Val)
	return temp
}