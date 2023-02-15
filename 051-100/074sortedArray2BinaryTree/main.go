package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	nums := []int{-10,-3,0,5,9}
	fmt.Println(sortedArrayToBST(nums))
}

//108 二叉搜索树 给中序遍历，恢复二叉树
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := len(nums)/2
	root := &TreeNode{
		Val: nums[index],
	}
	root.Left = sortedArrayToBST(nums[:index])
	root.Right = sortedArrayToBST(nums[index+1:])
	return root
}

