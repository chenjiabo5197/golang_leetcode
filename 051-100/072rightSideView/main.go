package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	n4 := &TreeNode{
		Val:   4,
	}
	n5 := &TreeNode{
		Val:   5,
	}
	n3 := &TreeNode{
		Val:   3,
		Right:n4,
	}
	n2 := &TreeNode{
		Val:   2,
		Right:n5,
	}
	n1 := &TreeNode{
		Val:   1,
		Left:n2,
		Right:n3,
	}
	fmt.Println(rightSideView(n1))
}

/*
	1、先判断当前层数节点数量，最右侧节点的索引是数量-1
	2、判断完一层，将每个节点从队列中弹出，然后将其左右子节点加入队列,严格按照先右再左的顺序
*/
// 199 二叉树的右视图
func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	queueSlice := make([]*TreeNode, 0)
	if root == nil {
		return result
	}
	queueSlice = append(queueSlice, root)
	for len(queueSlice) != 0 {
		tempNode := queueSlice[len(queueSlice)-1]
		result = append(result, tempNode.Val)
		length := len(queueSlice)  //保存本层切片长度
		for length > 0 {
			node := queueSlice[len(queueSlice)-1]
			queueSlice = append(queueSlice[:len(queueSlice)-1])
			length--
			if node.Right != nil {
				queueSlice = append([]*TreeNode{node.Right},queueSlice...)
			}
			if node.Left != nil {
				queueSlice = append([]*TreeNode{node.Left}, queueSlice...)
			}
		}
	}
	return result
}
