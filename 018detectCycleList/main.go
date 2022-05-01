package main

import "fmt"

/*
返回环形链表的入环节点
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l4 := &ListNode{
		Val:  -4,
		Next: nil,
	}
	l0 := &ListNode{
		Val:  0,
		Next: l4,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l0,
	}
	l3 := &ListNode{
		Val:  3,
		Next: l2,
	}
	l4.Next = l2
	fmt.Println(detectCycle(l3))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func detectCycle(head *ListNode) *ListNode {
//	tableAddr := make(map[*int]int)
//	tempNode := head
//	var result *ListNode
//	if head.Next == head {
//		return result
//	}
//	for {
//		if tempNode == nil {
//			break
//		}
//		if _, ok := tableAddr[&tempNode.Val]; ok {
//			result = &ListNode{
//				Val:  tempNode.Val,
//				Next: nil,
//			}
//			break
//		}
//		tableAddr[&tempNode.Val] = 1
//		tempNode = tempNode.Next
//	}
//	return result
//}

//func detectCycle(head *ListNode) *ListNode {
//	tableAddr := make(map[*int]bool)
//	tempNode := head
//	var result *ListNode
//	for {
//		if tempNode == nil {
//			break
//		}
//		fmt.Println(tempNode.Val)
//		if _, ok := tableAddr[&tempNode.Val]; ok {
//			return tempNode
//		}
//		tableAddr[&tempNode.Val] = true
//		tempNode = tempNode.Next
//
//	}
//	return result
//}

//使用快慢指针
/*
从头结点到圈的入口处是x，从入口处到第一次相遇的节点是y，从相遇节点到入口处节点是z，则一圈是y+z
慢指针走了x+y，快指针走了x+y+(y+z)n,其中n表示在圈中走了n圈
由于相遇  x+y = [x+y+(y+z)n]2 -> x = (y+z)n-y
即  若从相遇开始，有一个指针p1从头结点往后走，一个指针p2从相遇节点往后走，两者速度一样，则当p1走了x步时，p2走了(y+z)n-y步
    两者刚好在入口节点处相遇
 */
func detectCycle(head *ListNode) *ListNode {
	fastPoint := head
	slowPoint := head
	newPoint := head
	for {
		if fastPoint == nil || fastPoint.Next == nil {
			return nil  //证明无环
		}
		slowPoint = slowPoint.Next
		fastPoint = fastPoint.Next.Next
		if fastPoint == slowPoint {
			break
		}
	}
	for {
		if newPoint == slowPoint {
			return newPoint
		}
		slowPoint = slowPoint.Next
		newPoint = newPoint.Next
	}
}
