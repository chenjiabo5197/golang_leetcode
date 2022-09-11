package main

import "fmt"

/*
翻转链表
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l1,
	}
	//l3 := &ListNode{
	//	Val:  3,
	//	Next: l2,
	//}
	newHead := reverseList(l2)
	for {
		if newHead == nil {
			break
		}
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func reverseList(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//	if head.Next == nil {
//		return head
//	}
//	tempHeadNext := &ListNode{
//		Val:head.Val,
//		Next:nil,
//	}
//	newHead := &ListNode{
//		Val:head.Next.Val,
//		Next:tempHeadNext,
//	}
//	if head.Next.Next == nil {
//		return newHead
//	}
//	nowPoint := &ListNode{
//		Val:head.Next.Next.Val,
//		Next:head.Next.Next.Next,
//	}
//	for {
//		if nowPoint == nil {
//			break
//		}
//		tempHeadNext = newHead
//		newHead = nowPoint
//		nowPoint = nowPoint.Next
//		newHead.Next = tempHeadNext
//	}
//	return newHead
//}


//使用递归的方式
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	//fmt.Println(newHead)
	//fmt.Println(head)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
