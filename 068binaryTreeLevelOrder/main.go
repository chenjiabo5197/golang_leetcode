package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	n8 := &TreeNode{
		Val:   8,
	}
	n6 := &TreeNode{
		Val:   6,
	}
	n11 := &TreeNode{
		Val:   1,
	}
	n5 := &TreeNode{
		Val:   5,
	}
	n12 := &TreeNode{
		Val:   -1,
		Right:n8,
	}
	n3 := &TreeNode{
		Val:   3,
		Right:n6,
	}
	n13 := &TreeNode{
		Val:   1,
		Left:n5,
		Right:n11,
	}
	n4 := &TreeNode{
		Val:   4,
		Left:n3,
		Right:n12,
	}
	n2 := &TreeNode{
		Val:   2,
		Left:n13,
	}
	n0 := &TreeNode{
		Val:   0,
		Left:n2,
		Right:n4,
	}
	fmt.Println(levelOrder(n0))
	fmt.Println(zigzagLevelOrder(n0))
}

// 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	var temp [][]int
	if root == nil {
		return temp
	}
	var addrSlice []map[*TreeNode]int   //储存节点和层数的对应关系,因为map无序，所以用切片来保证顺序
	addrSlice2 := preorderTraversal(root, 0, addrSlice)
	var maxLevel int
	for i:= range addrSlice2 {
		for _, value := range addrSlice2[i] {
			if maxLevel < value {
				maxLevel = value
			}
		}
	}
	temp = make([][]int, maxLevel+1)
	for i:= range addrSlice2 {
		for key, value := range addrSlice2[i] {
			temp[value] = append(temp[value], key.Val)
		}
	}
	return temp
}

func preorderTraversal(root *TreeNode, level int, addrSlice []map[*TreeNode]int) []map[*TreeNode]int {
	if root == nil {
		return addrSlice
	}
	tempMap := make(map[*TreeNode]int)
	tempMap[root] = level
	addrSlice = append(addrSlice, tempMap)
	addrSlice = preorderTraversal(root.Left, level+1, addrSlice)
	addrSlice = preorderTraversal(root.Right, level+1, addrSlice)
	return addrSlice
}

// 锯齿形遍历  上层先从左到右，下一层从右到左
func zigzagLevelOrder(root *TreeNode) [][]int {
	var temp [][]int
	if root == nil {
		return temp
	}
	var addrSlice []map[*TreeNode]int   //储存节点和层数的对应关系,因为map无序，所以用切片来保证顺序
	addrSlice2 := preorderTraversal(root, 0, addrSlice)
	var maxLevel int
	for i:= range addrSlice2 {
		for _, value := range addrSlice2[i] {
			if maxLevel < value {
				maxLevel = value
			}
		}
	}
	temp = make([][]int, maxLevel+1)
	for i:= range addrSlice2 {
		for key, value := range addrSlice2[i] {
			temp[value] = append(temp[value], key.Val)
		}
	}
	for i := range temp {
		if i%2 != 0 {
			reverseSlice(temp[i])
		}
	}
	return temp
}

func reverseSlice (data []int) {
	var tempSlice []int
	for i:=0;i<len(data);i++ {
		tempSlice = append(tempSlice, data[i])
	}
	fmt.Println(tempSlice)
	for i:=0;i<len(data);i++ {
		data[i] = tempSlice[len(data)-i-1]
	}
}
