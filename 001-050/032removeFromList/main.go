package main

import "fmt"

/*
移除链表元素
 */

func main() {
	l61 := &ListNode{
		Val:  6,
		Next: nil,
	}
	l5 := &ListNode{
		Val:  5,
		Next: l61,
	}
	l4 := &ListNode{
		Val:  4,
		Next: l5,
	}
	l3 := &ListNode{
		Val:  3,
		Next: l4,
	}
	l62 := &ListNode{
		Val:  6,
		Next: l3,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l62,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}
	head := removeElements(l1, 6)
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

type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{
		Next:head,
	}
	preNode := newHead
	for{
		if preNode.Next == nil {
			break
		}
		if preNode.Next.Val == val {
			preNode.Next = preNode.Next.Next
		}else {
			preNode = preNode.Next
		}
	}
	return newHead.Next
}
