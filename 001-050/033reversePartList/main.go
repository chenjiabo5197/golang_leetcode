package main

import "fmt"

/*
k个一组翻转链表
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
	head := reverseKGroup(l1, 2)
	for {
		if head == nil {
			break
		}
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	newHead := &ListNode{
		Val:   -1,
		Next: head,
	}
	preNode := newHead  //指向要翻转区域的前一个节点
	endPoint := newHead  //指向要翻转链表的尾结点
	for endPoint != nil {
		for times:=0;times<k;times++ {
			if endPoint == nil {
				break
			}
			endPoint = endPoint.Next
		}
		if endPoint != nil {
			tempNode := endPoint.Next
			newEnd := preNode.Next
			reverseList(preNode.Next, endPoint)  // 翻转指定范围的节点
			preNode.Next = endPoint  // 将翻转完成之后的链表连接到原来链表中
			newEnd.Next = tempNode
			preNode = newEnd    //更新preNode和endPoint,开始下一次循环
			endPoint = newEnd
		}
	}
	return newHead.Next
}

func reverseList(head, rear *ListNode) *ListNode {
	if head == rear  {
		return head
	}
	nowHead := reverseList(head.Next, rear)
	head.Next.Next = head
	head.Next = nil
	return nowHead
}
