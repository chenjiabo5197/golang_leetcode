package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l9 := &ListNode{
		Val: 9,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l9,
	}
	l5 := &ListNode{
		Val:  5,
		Next: l1,
	}
	l4 := &ListNode{
		Val:  4,
		Next: l5,
	}
	val := 5
	fmt.Println(deleteNode(l4, val))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//剑指 Offer 18. 删除链表的节点
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{}
	newHead.Next = head
	preNode := newHead
	curNode := head
	for curNode.Val != val {
		curNode = curNode.Next
		preNode = preNode.Next
	}
	preNode.Next = curNode.Next
	return newHead.Next
}
