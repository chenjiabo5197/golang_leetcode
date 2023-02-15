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
	n7 := &TreeNode{
		Val:   7,
	}
	n8 := &TreeNode{
		Val:   8,
	}
	n0 := &TreeNode{
		Val:   0,
	}
	n2 := &TreeNode{
		Val:   2,
		Left:n7,
		Right:n4,
	}
	n6 := &TreeNode{
		Val:   6,
	}
	n1 := &TreeNode{
		Val:   1,
		Left:n0,
		Right:n8,
	}
	n5 := &TreeNode{
		Val:   5,
		Left:n6,
		Right:n2,
	}
	n3 := &TreeNode{
		Val:   3,
		Left:n5,
		Right:n1,
	}

	fmt.Println(lowestCommonAncestor(n3, n5, n1))
}

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	nodeMap := make(map[*TreeNode][]*TreeNode, 0)
// 	preOrder(root, root, nodeMap)
// 	pSlice := nodeMap[p]
// 	qSlice := nodeMap[q]
// 	for i:=len(pSlice)-1;i>=0;i-- {
// 		for j:=len(qSlice)-1;j>=0;j-- {
// 			if pSlice[i] == qSlice[j] {
// 				return pSlice[i]
// 			}
// 		}
// 	}
// 	return nil
// }

// func preOrder(root, father *TreeNode, nodeMap map[*TreeNode][]*TreeNode) map[*TreeNode][]*TreeNode {
// 	if root == nil {
// 		return nodeMap
// 	}
// 	nodeMap[root] = append(nodeMap[root], nodeMap[father]...)
// 	nodeMap[root] = append(nodeMap[root], root)
// 	nodeMap = preOrder(root.Left, root, nodeMap)
// 	nodeMap = preOrder(root.Right, root, nodeMap)
// 	return nodeMap
// }

// 236 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p {
		return p
	}
	if root == q {
		return q
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right == nil {
		return nil
	}else if left == nil {
		return right
	}else if right == nil {
		return left
	}else {
		return root
	}
}
