package main

import "fmt"

/*
奇偶链表
 */

func main() {
	l5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l4 := &ListNode{
		Val:  4,
		Next: l5,
	}
	l3 := &ListNode{
		Val:  3,
		Next: l4,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l3,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}
	head := oddEvenList(l1)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
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
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newEvenHead := &ListNode{   //保存偶数节点头结点的地址
		Val:-1,
		Next:head.Next,
	}
	preOddNode := head
	preEvenNode := head.Next
	curNode := head.Next.Next
	flag := 1
	for curNode != nil {
		if flag % 2 == 0 {
			preEvenNode.Next = curNode
			preEvenNode = preEvenNode.Next
		}else {
			preOddNode.Next = curNode
			preOddNode = preOddNode.Next
		}
		curNode = curNode.Next
		flag++
	}
	preEvenNode.Next = nil
	preOddNode.Next = newEvenHead.Next
	return head
}
