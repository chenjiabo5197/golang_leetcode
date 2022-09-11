package main

import "fmt"

/*
合并k个有序数组
 */

func main() {
	//l6 := &ListNode{
	//	Val:  6,
	//	Next: nil,
	//}
	//l2 := &ListNode{
	//	Val:  2,
	//	Next: l6,
	//}
	//l4 := &ListNode{
	//	Val:  4,
	//	Next: nil,
	//}
	//l3 := &ListNode{
	//	Val:  3,
	//	Next: l4,
	//}
	//l1 := &ListNode{
	//	Val:  1,
	//	Next: l3,
	//}
	//l5 := &ListNode{
	//	Val:  5,
	//	Next: nil,
	//}
	l42 := &ListNode{
		Val:  -1,
		Next: nil,
	}
	l12 := &ListNode{
		Val:  2,
		Next: nil,
	}
	var temp *ListNode
	data := []*ListNode{l12, temp, l42}
	head := mergeKLists(data)
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
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	head := &ListNode{
		Val:  -1,
		Next: nil,
	}
	preNode := head
	index := -1
	for len(lists)>0 {
		for i:=0;i<len(lists); {
			if lists[i] == nil {   //删除已经为空的链表
				lists = append(lists[:i], lists[i+1:]...)
				continue
			}
			if index == -1 {   //给index赋予每次循环开始的值
				index = i
			}
			if lists[i].Val < lists[index].Val {
				index = i
			}
			i++
		}
		if index != -1 {
			if head.Next == nil {
				head.Next = lists[index]
				preNode = preNode.Next
			}else {
				preNode.Next = lists[index]
				preNode = preNode.Next
			}
			lists[index] = lists[index].Next
			if lists[index] == nil {   //删除已经为空的链表
				lists = append(lists[:index], lists[index+1:]...)
			}
			index = -1
		}
	}
	return head.Next
}
