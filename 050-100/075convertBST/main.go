package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	n8 := &TreeNode{
		Val:   8,
	}
	n3 := &TreeNode{
		Val:   3,
	}
	n7 := &TreeNode{
		Val:   7,
		Right:n8,
	}
	n5 := &TreeNode{
		Val:   5,
	}
	n2 := &TreeNode{
		Val:   2,
		Right:n3,
	}
	n0 := &TreeNode{
		Val:   0,
	}
	n6 := &TreeNode{
		Val:   6,
		Left:n5,
		Right:n7,
	}
	n1 := &TreeNode{
		Val:   1,
		Left:n0,
		Right:n2,
	}
	n4 := &TreeNode{
		Val:   4,
		Left:n1,
		Right:n6,
	}
	fmt.Println(convertBST(n4))
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	addNodeValue(root, 0)
	return root
}

// preNode代表父节点，nowNode是当前节点，rightValue用于当前父节点的左节点的右节点的值要加的数
func addNodeValue(node *TreeNode, rightValue int) int {
	if node.Right != nil {
		rightValue = addNodeValue(node.Right, rightValue)
	}
	node.Val += rightValue
	leftValue := -10000
	if node.Left != nil {
		leftValue = addNodeValue(node.Left, node.Val)
	}
	if leftValue != -10000 {
		return leftValue
	}
	return node.Val
}