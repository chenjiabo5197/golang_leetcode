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
	fmt.Println(reverseList(l4))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	preNode := &ListNode{}
	preNode.Next = head
	curNode := head
	for curNode.Next != nil {
		curNode = curNode.Next
		preNode = preNode.Next
	}
}
