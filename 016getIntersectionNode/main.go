package main

import (
	"fmt"
)

/*
获得相交链表的节点
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l4,
	}
	l11 := &ListNode{
		Val:  1,
		Next: l2,
	}
	l9 := &ListNode{
		Val:  9,
		Next: l11,
	}
	l12 := &ListNode{
		Val:  1,
		Next: l9,
	}
	l3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	fmt.Println(getIntersectionNode(l12, l3))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	addrAList := make(map[*int]int)
//	addrBList := make(map[*int]int)
//	tempA := headA
//	tempB := headB
//	result := &ListNode{}
//	flag := true
//	for {
//		if tempA == nil && tempB == nil{
//			break
//		}
//		if tempA != nil {
//			addrAList[&tempA.Val] = 1
//			if _, ok := addrBList[&tempA.Val]; ok {
//				result.Val = tempA.Val
//				flag = false
//				break
//			}
//			tempA = tempA.Next
//		}
//		if tempB != nil {
//			addrBList[&tempB.Val] = 1
//			if _, ok := addrAList[&tempB.Val]; ok {
//				result.Val = tempB.Val
//				flag = false
//				break
//			}
//			tempB = tempB.Next
//		}
//	}
//	if flag {
//		return  nil
//	}
//	return result
//}

//使用两个指针，不另外开辟空间，最多循环两次即可找到相同的节点或者nil
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pointA := headA
	pointB := headB
	for {
		if pointA == pointB {
			break
		}
		if pointA == nil {
			pointA = headB
		}else {
			pointA = pointA.Next
		}
		if pointB == nil {
			pointB = headA
		}else {
			pointB = pointB.Next
		}
	}
	return pointA
}
