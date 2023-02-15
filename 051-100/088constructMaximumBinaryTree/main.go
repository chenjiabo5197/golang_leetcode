package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	result := []int{3,2,1,6,0,5}
	fmt.Println(constructMaximumBinaryTree(result))
}

// 654. 最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxIndex := getMaxValueIndex(nums)
	root := &TreeNode{
		Val:   nums[maxIndex],
	}
	root.Left = constructMaximumBinaryTree(nums[0:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	return root
}

func getMaxValueIndex(nums []int) int {
	maxIndex := 0
	for index := range nums {
		if nums[index] > nums[maxIndex] {
			maxIndex = index
		}
	}
	return maxIndex
}
