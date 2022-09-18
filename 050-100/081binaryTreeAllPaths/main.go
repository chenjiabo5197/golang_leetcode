package main

import (
	"fmt"
	"strconv"
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
	fmt.Println(binaryTreePaths(test()))
}

// 257 二叉树的所有路径   ["1->2->5","1->3"]
func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	if root == nil {
		return result
	}
	result = append(result, strconv.Itoa(root.Val))
	result = getPaths(root, result)
	return result
}

func getPaths(node *TreeNode, data []string) []string {
	temp := make([]string, len(data))
	copy(temp, data)
	if node == nil {
		return data
	}
	var data1, data2 []string
	if node.Left != nil {  //更新当前切片
		for index := range temp {
			temp[index] = temp[index] + "->" + strconv.Itoa(node.Left.Val)
		}
		data1 = getPaths(node.Left, temp)
	}
	if node.Right != nil {
		for index := range data {
			data[index] = data[index] + "->" + strconv.Itoa(node.Right.Val)
		}
		data2 = getPaths(node.Right, data)
	}
	if len(data1) == 0 && len(data2) == 0 {
		return temp
	}
	data2 = append(data2, data1...)
	return data2
}

