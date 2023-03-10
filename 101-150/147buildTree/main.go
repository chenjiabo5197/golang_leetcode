package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	fmt.Println(buildTree(preorder, inorder))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//剑指 Offer 07. 重建二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	if len(preorder) == 1 {
		return root
	}
	inorderIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			inorderIndex = i
			break
		}
	}
	if inorderIndex == 0 { // 无左子树
		root.Right = buildTree(preorder[1:], inorder[1:])
	} else if inorderIndex == len(inorder)-1 { // 无右子树
		root.Left = buildTree(preorder[1:], inorder[0:inorderIndex])
	} else {
		preorderIndex := 0
		for i := 0; i < len(preorder); i++ {
			for j := 0; j < inorderIndex; j++ {
				if preorder[i] == inorder[j] {
					preorderIndex = max(preorderIndex, i)
				}
			}
		}
		root.Left = buildTree(preorder[1:preorderIndex+1], inorder[0:inorderIndex])
		root.Right = buildTree(preorder[preorderIndex+1:], inorder[inorderIndex+1:])
	}
	return root
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
