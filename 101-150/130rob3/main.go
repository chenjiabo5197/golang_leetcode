package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	n1 := &TreeNode{
		Val: 1,
	}
	n3 := &TreeNode{
		Val: 3,
	}
	n12 := &TreeNode{
		Val: 1,
	}
	n5 := &TreeNode{
		Val:   5,
		Right: n1,
	}
	n4 := &TreeNode{
		Val:   4,
		Left:  n12,
		Right: n3,
	}
	n32 := &TreeNode{
		Val:   3,
		Left:  n4,
		Right: n5,
	}
	fmt.Println(rob(n32))
}

/*
	优化，将子节点的值进行储存
*/

var (
	nodeValueMap = make(map[*TreeNode]int, 0)
)

func rob(root *TreeNode) int {
	nodeValue, ok := nodeValueMap[root]
	if ok {
		return nodeValue
	}
	if root == nil {
		return 0
	}
	value := root.Val
	if root.Left != nil {
		value += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		value += rob(root.Right.Left) + rob(root.Right.Right)
	}
	nodeValueMap[root] = max(value, rob(root.Left)+rob(root.Right))
	return nodeValueMap[root]
}

/**
 * 递归，爷爷节点的值加上四个孙子节点的值和两个儿子节点的值进行比较，超时
 */
//func rob(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	value := root.Val
//	if root.Left != nil {
//		value += rob(root.Left.Left) + rob(root.Left.Right)
//	}
//	if root.Right != nil {
//		value += rob(root.Right.Left) + rob(root.Right.Right)
//	}
//	return max(value, rob(root.Left)+rob(root.Right))
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
