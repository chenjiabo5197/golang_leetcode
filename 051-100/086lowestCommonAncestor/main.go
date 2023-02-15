package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func test() (*TreeNode, *TreeNode, *TreeNode) {
	n5 := &TreeNode{
		Val:   5,
	}
	n3 := &TreeNode{
		Val:   3,
	}
	n9 := &TreeNode{
		Val:   9,
	}
	n7 := &TreeNode{
		Val:   7,
	}
	n4 := &TreeNode{
		Val:   4,
		Left:n3,
		Right:n5,
	}
	n0 := &TreeNode{
		Val:   0,
	}
	n8 := &TreeNode{
		Val:   8,
		Left:n7,
		Right:n9,
	}
	n2 := &TreeNode{
		Val:   2,
		Left:n0,
		Right:n4,
	}
	n6 := &TreeNode{
		Val:   6,
		Left:n2,
		Right:n8,
	}
	return n6, n2, n8
}

func main()  {
	fmt.Println(lowestCommonAncestor(test()))
}

// 235 二叉搜索树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	if isMiddle(root, p, q){
		return root
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}else {
		return lowestCommonAncestor(root.Right, p, q)
	}
}

func isMiddle(root, p, q *TreeNode) bool {
	if (root.Val > p.Val && root.Val < q.Val) || (root.Val < p.Val && root.Val > q.Val) {
		return true
	}
	return false
}
