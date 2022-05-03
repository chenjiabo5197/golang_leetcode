package main

import "fmt"

/*
判断是否回文链表
 */

func main() {
	l11 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l21 := &ListNode{
		Val:  2,
		Next: l11,
	}
	l22 := &ListNode{
		Val:  2,
		Next: l21,
	}
	l12 := &ListNode{
		Val:  1,
		Next: l22,
	}
	fmt.Println(isPalindrome(l12))
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
func isPalindrome(head *ListNode) bool {
	addr := make([]*int, 0)
	for head != nil {
		addr = append(addr, &head.Val)
		head = head.Next
	}
	half := len(addr) / 2
	for i:=0;i<half;i++ {
		if *addr[i] != *addr[len(addr)-i-1] {
			return false
		}
	}
	return true
}