package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	n1 := &TreeNode{
		Val:   1,
	}
	n5 := &TreeNode{
		Val:   5,
	}
	n2 := &TreeNode{
		Val:   2,
	}
	n7 := &TreeNode{
		Val:   7,
	}
	n4 := &TreeNode{
		Val:   4,
		Left:n5,
		Right:n1,
	}
	n13 := &TreeNode{
		Val:   13,
	}
	n11 := &TreeNode{
		Val:   11,
		Left:n7,
		Right:n2,
	}
	n8 := &TreeNode{
		Val:   8,
		Left:n13,
		Right:n4,
	}
	n42 := &TreeNode{
		Val:   4,
		Left:n11,
	}
	n52 := &TreeNode{
		Val:   5,
		Left:n42,
		Right:n8,
	}
	targetSum := 22
	fmt.Println(pathSum(n52, targetSum))
}

// 路径总和
func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	var temp []int
	if root == nil {
		return result
	}
	temp = append(temp, root.Val)
	return countSum(root, root.Val, targetSum, temp, result)
}

func countSum(root *TreeNode, tempSum, target int, tempSlice []int, result [][]int) [][]int{
	if root.Left == nil && root.Right == nil && tempSum == target {
		result = append(result, tempSlice)
		return result
	}
	if root.Left != nil{
		tempSlice1 := make([]int, len(tempSlice))
		copy(tempSlice1, tempSlice)
		tempSlice1 = append(tempSlice1, root.Left.Val)
		result = countSum(root.Left, tempSum+root.Left.Val, target, tempSlice1, result)
	}
	if root.Right != nil{
		tempSlice2 := make([]int, len(tempSlice))
		copy(tempSlice2, tempSlice)    // 不改成这种拷贝方式会一直出错
		tempSlice2 = append(tempSlice2, root.Right.Val)   // 这里会将已经append的结果修改掉，应该是切片是引用类型的原因
		result = countSum(root.Right, tempSum+root.Right.Val, target, tempSlice2, result)
	}
	return result
}
