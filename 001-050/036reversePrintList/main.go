package main

import "fmt"

/*
翻转打印链表
 */

func main() {
	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l3 := &ListNode{
		Val:  3,
		Next: l2,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l3,
	}
	fmt.Println(reversePrint(l1))
}

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	newHead := reverseList(head)
	result := make([]int, 0)
	for newHead != nil {
		result = append(result, newHead.Val)
		newHead = newHead.Next
	}
	return result
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curNode := reverseList(head.Next)  //curNode不做改动，所以一直指向新头结点
	head.Next.Next = head
	head.Next = nil
	return curNode
}
