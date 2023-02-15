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
	n49 := &TreeNode{
		Val:   49,
	}
	n12 := &TreeNode{
		Val:   12,
	}
	n48 := &TreeNode{
		Val:   48,
		Left:n12,
		Right:n49,
	}
	n0 := &TreeNode{
		Val:   0,
	}
	n1 := &TreeNode{
		Val:   1,
		Left:n0,
		Right:n48,
	}
	return n1
}

func main()  {
	fmt.Println(getMinimumDifference(test()))
}

// 根节点需要和左子树最大节点的值，右子树最小节点的值进行比较
// 530. 二叉搜索树的最小绝对差
func getMinimumDifference(root *TreeNode) int {
	minValue := 10000
	if root == nil {
		return minValue
	}
	if root.Left != nil {
		node := getMaxValueNode(root.Left)
		if root.Val - node.Val < minValue {
			minValue = root.Val - node.Val
		}
	}
	if root.Right != nil {
		node := getMinValueNode(root.Right)
		if node.Val - root.Val < minValue {
			minValue = node.Val - root.Val
		}
	}
	leftMin := getMinimumDifference(root.Left)
	rightMin := getMinimumDifference(root.Right)
	return int(math.Min(math.Min(float64(leftMin), float64(rightMin)), float64(minValue)))
}

// 获取最小值节点
func getMinValueNode(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

// 获取最大值节点
func getMaxValueNode(root *TreeNode) *TreeNode {
	for root.Right != nil {
		root = root.Right
	}
	return root
}
