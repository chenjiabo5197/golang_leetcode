package main

import "fmt"

/*
得到倒数第k个节点
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
	fmt.Println(getKthFromEnd(l1, 2))
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
//func getKthFromEnd(head *ListNode, k int) *ListNode {
//	addrVal := make([]*int, 0)
//	addrNext := make([]**ListNode, 0)
//	for head != nil {
//		addrVal = append(addrVal, &head.Val)
//		addrNext = append(addrNext, &head.Next)
//		head = head.Next
//	}
//	return &ListNode{
//		Val:  *addrVal[len(addrVal)-k],
//		Next: *addrNext[len(addrNext)-k],
//	}
//}

//使用两个指针，遍历一次链表即可
func getKthFromEnd(head *ListNode, k int) *ListNode {
	temp := &ListNode{
		Val:  -1,
		Next: head,
	}
	behindPoint := temp
	frontPoint := temp
	for i:=0;i<k;i++ {
		frontPoint = frontPoint.Next
	}
	for frontPoint != nil {
		frontPoint = frontPoint.Next
		behindPoint = behindPoint.Next
	}
	return behindPoint
}
