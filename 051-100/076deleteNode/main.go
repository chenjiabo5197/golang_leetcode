package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	n7 := &TreeNode{
		Val:   7,
	}
	n4 := &TreeNode{
		Val:   4,
	}
	n2 := &TreeNode{
		Val:   2,
	}
	n6 := &TreeNode{
		Val:   6,
		Right:n7,
	}
	n3 := &TreeNode{
		Val:   3,
		Left:n2,
		Right:n4,
	}
	n5 := &TreeNode{
		Val:   5,
		Left:n3,
		Right:n6,
	}
	key := 3
	fmt.Println(deleteNode(n5, key))
}
/*
迭代函数
	1、节点是叶子节点					返回nil
	2、节点只有左子树或者右子树		返回左/右子树
	3、节点左右子树都有				用左子树的最右侧节点的值替代，再删除左子树最右侧节点
 */
// 450 删除二叉搜索树中的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}else {
		if root.Left == nil && root.Right == nil {
			return nil
		}else if root.Left == nil {
			return root.Right
		}else if root.Right == nil {
			return root.Left
		}else {
			node := getLeftBiggestNode(root.Left)
			root.Val = node.Val
			root.Left = deleteNode(root.Left, node.Val)
		}
	}
	return root
}

func getLeftBiggestNode(node *TreeNode) *TreeNode {
	for node.Right != nil {
		node = node.Right
	}
	return node
}
