package main

import "fmt"

/*
合并两个有序链表
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l5 := ListNode{
		Val:  5,
		Next: nil,
	}
	l3 := ListNode{
		Val:  3,
		Next: &l5,
	}
	l1 := ListNode{
		Val:  1,
		Next: &l3,
	}
	l6 := ListNode{
		Val:  6,
		Next: nil,
	}
	l4 := ListNode{
		Val:  4,
		Next: &l6,
	}
	l2 := ListNode{
		Val:  2,
		Next: &l4,
	}
	l7 := mergeTwoLists(&l1, &l2)
	for {
		fmt.Println(l7.Val)
		l7 = l7.Next
		if l7 == nil {
			break
		}
	}

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		if list2 == nil {
			return nil
		}
		return list2
	}
	if list2 == nil {
		return list1
	}
	result := &ListNode{}
	head := &ListNode{Next:result}
	for {
		if list1 == nil || list2 == nil {
			break
		}
		if list1.Val < list2.Val {
			temp := &ListNode{
				Val:  list1.Val,
				Next: nil,
			}
			result.Next = temp
			result = result.Next
			list1 = list1.Next
		}else {
			temp := &ListNode{
				Val:  list2.Val,
				Next: nil,
			}
			result.Next = temp
			result = result.Next
			list2 = list2.Next
		}
	}
	if list1 == nil {
		result.Next = list2
	}
	if list2 == nil {
		result.Next = list1
	}

	return head.Next.Next

}

