package main

import "fmt"

/*
分隔链表
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l21 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l5 := &ListNode{
		Val:  5,
		Next: l21,
	}
	l22 := &ListNode{
		Val:  2,
		Next: l5,
	}
	l3 := &ListNode{
		Val:  3,
		Next: l22,
	}
	l4 := &ListNode{
		Val:  4,
		Next: l3,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l4,
	}
	newList := partition(l1, 3)
	for {
		if newList == nil {
			break
		}
		fmt.Println(newList.Val)
		newList = newList.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 由于需要节点的相对位置不变，所以放弃该方法
//func partition(head *ListNode, x int) *ListNode {
//	if head.Next == nil {
//		return head
//	}
//	tempNode := head
//	tempNext := head.Next
//	for {
//		if tempNext == nil {
//			break
//		}
//		if tempNode.Val >= x {
//			for {
//				if tempNext == nil {
//					break
//				}
//				if tempNext.Val < x {
//					tempVal := tempNext.Val
//					tempNext.Val = tempNode.Val
//					tempNode.Val = tempVal
//					break
//				}
//				tempNext = tempNext.Next
//			}
//		}
//		tempNode = tempNode.Next
//		tempNext = tempNext.Next
//	}
//	return head
//}

//func partition(head *ListNode, x int) *ListNode {
//	var smallHead, smallRear, bigHead, bigRear *ListNode
//	temp := head
//	for {
//		if temp == nil {
//			break
//		}
//		if temp.Val < x {
//			if smallHead == nil{
//				smallHead = &ListNode{
//					Val:  temp.Val,
//					Next: nil,
//				}
//			}else if smallHead != nil && smallRear == nil {
//				smallRear = &ListNode{
//					Val:  temp.Val,
//					Next: nil,
//				}
//				smallHead.Next = smallRear
//			}else {
//				tempNode := &ListNode{
//					Val:  temp.Val,
//					Next: nil,
//				}
//				smallRear.Next = tempNode
//				smallRear = smallRear.Next
//			}
//		}else {
//			if bigHead == nil{
//				bigHead = &ListNode{
//					Val:  temp.Val,
//					Next: nil,
//				}
//			}else if bigHead != nil && bigRear == nil {
//				bigRear = &ListNode{
//					Val:  temp.Val,
//					Next: nil,
//				}
//				bigHead.Next = bigRear
//			}else {
//				tempNode := &ListNode{
//					Val:  temp.Val,
//					Next: nil,
//				}
//				bigRear.Next = tempNode
//				bigRear = bigRear.Next
//			}
//		}
//		temp = temp.Next
//	}
//	if smallHead == nil {
//		return bigHead
//	}
//	if smallRear == nil {
//		smallHead.Next = bigHead
//	}else {
//		smallRear.Next = bigHead
//	}
//	return smallHead
//}


//采用虚拟头结点，代码简化
func partition(head *ListNode, x int) *ListNode {
	bigHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	bigRear := bigHead
	smallHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	smallRear := smallHead
	for {
		if head == nil {
			break
		}
		if head.Val < x {
			smallRear.Next = head
			smallRear = smallRear.Next
		}else {
			bigRear.Next = head
			bigRear = bigRear.Next
		}
		head = head.Next
	}
	bigRear.Next =nil
	smallRear.Next =bigHead.Next
	return smallHead.Next
}
