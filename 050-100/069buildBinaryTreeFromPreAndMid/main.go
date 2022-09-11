package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	preOrder := []int{3,9,20,15,7}
	inOrder := []int{9,3,15,20,7}
	fmt.Println(buildTree(preOrder, inOrder))
}

//会超时,但容易理解
//func buildTree(preorder []int, inorder []int) *TreeNode {
//	if len(preorder) == 0 || len(inorder) == 0 {
//		return nil
//	}
//	root := &TreeNode{
//		Val: preorder[0],
//	}
//	var index1 int   //找到在中序中根节点的位置
//	for index1 = range inorder {
//		if inorder[index1] == preorder[0] {
//			break
//		}
//	}
//	var index2 int   //找到在前序中左边全部的位置
//	if index1 != 0 {
//		index2 = getPreOrderIndex(preorder, inorder[:index1])
//	}
//	root.Left = buildTree(preorder[1:index2+1], inorder[0:index1+1])
//	root.Right = buildTree(preorder[index2+1:], inorder[index1+1:])
//	return root
//}
//
//func getPreOrderIndex(preOrder, inOrder []int) int {
//	var maxIndex int
//	for i := range inOrder{
//		for j := range preOrder {
//			if inOrder[i] == preOrder[j] && maxIndex<j {
//				maxIndex = j
//				break
//			}
//		}
//	}
//	return maxIndex
//}

// 根据前序和中序构建二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	//建立一个中序遍历key和位置对应的map
	indexMap := make(map[int]int)
	for i:=0;i<len(inorder);i++ {
		indexMap[inorder[i]] = i
	}
	for i:=1;i<len(preorder);i++ {
		node := &TreeNode{Val: preorder[i]}
		build(indexMap, root, node)
	}
	return root
}

func build(data map[int]int, root, node *TreeNode) {

	for root != node{
		if data[node.Val] < data[root.Val] {
			if root.Left == nil {
				root.Left = node
			}
			root = root.Left
		}
		if data[node.Val] > data[root.Val]{
			if root.Right == nil {
				root.Right = node
			}
			root = root.Right
		}
	}
}
