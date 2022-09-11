package main

import "fmt"

/*
翻转区间链表
 */

type ListNode struct {
	Val int
	Next *ListNode
}

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
	head := reverseBetween(l1, 2, 4)
	for {
		if head == nil {
			break
		}
		fmt.Println(head.Val)
		head = head.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func reverseBetween(head *ListNode, left int, right int) *ListNode {
//	if left == right {
//		return head
//	}
//	index := left
//	temp := head.Next
//	leftRear := head
//	for {
//		if left <= 2 {
//			break
//		}
//		temp = temp.Next
//		leftRear = leftRear.Next
//		left --
//	}
//	if head == nil {
//		return nil
//	}
//	if head.Next == nil {
//		return head
//	}
//	tempHeadNext := &ListNode{
//		Val:leftRear.Val,
//		Next:nil,
//	}
//	newHead := &ListNode{
//		Val:leftRear.Next.Val,
//		Next:tempHeadNext,
//	}
//	if leftRear.Next.Next == nil {
//		return newHead
//	}
//	nowPoint := &ListNode{
//		Val:leftRear.Next.Next.Val,
//		Next:leftRear.Next.Next.Next,
//	}
//	for {
//		index++
//		if index == right {
//			break
//		}
//		tempHeadNext = newHead
//		newHead = nowPoint
//		nowPoint = nowPoint.Next
//		newHead.Next = tempHeadNext
//	}
//	leftRear.Next = newHead
//	temp = head
//	for {
//		if temp.Next == nil {
//			temp.Next = nowPoint
//			break
//		}
//		temp = temp.Next
//	}
//	return head
//}

//添加虚拟节点，因为需要一个指针，指向当前节点的前一个节点
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	virtualNode := &ListNode{
		Val:  -1,
		Next: head,
	}
	prePoint := virtualNode
	curPoint := head
	//先寻找到要开始翻转的节点,此时curPoint指向要开始翻转的节点
	for i:=1;i<left;i++ {
		prePoint = prePoint.Next
		curPoint = curPoint.Next
	}
	//从left翻转到right
	for i:=left;i<right;i++ {
		tempPoint := curPoint.Next
		curPoint.Next = tempPoint.Next
		tempPoint.Next = prePoint.Next
		prePoint.Next = tempPoint
	}
	return virtualNode.Next
}
