package main

import "fmt"

/*
深拷贝链表
 */

type Node struct {
	Val int
	Next *Node
	Random *Node
}

func main() {
	n7 := &Node{
		Val:    7,
		Next:   nil,
		Random: nil,
	}
	n13 := &Node{
		Val:    13,
		Next:   nil,
		Random: nil,
	}
	n11 := &Node{
		Val:    11,
		Next:   nil,
		Random: nil,
	}
	n10 := &Node{
		Val:    10,
		Next:   nil,
		Random: nil,
	}
	n1 := &Node{
		Val:    1,
		Next:   nil,
		Random: nil,
	}
	n7.Next = n13
	n13.Next = n11
	n13.Random = n7
	n11.Next = n10
	n11.Random = n1
	n10.Next = n1
	n10.Random = n11
	n1.Random = n7
	head := copyRandomList(n7)
	for {
		if head == nil {
			break
		}
		fmt.Println(head.Val)
		head = head.Next
	}
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

//遍历两次，第一次建立拷贝的节点，第二次将指针指向位置
func copyRandomList(head *Node) *Node {
	curNode := head
	tableNode := make(map[*Node]*Node)
	for {
		if curNode == nil {
			break
		}
		tempNode := &Node{
			Val:    curNode.Val,
			Next:   nil,
			Random: nil,
		}
		tableNode[curNode] = tempNode
		curNode = curNode.Next
	}
	curNode = head
	for {
		if curNode == nil {
			break
		}
		//通过从map中取值将指针指向校正,先找到新链表中当前节点
		newCurNode := tableNode[curNode]

		oldCurNext := curNode.Next
		newCurNext := tableNode[oldCurNext]
		newCurNode.Next = newCurNext

		oldCurRandom := curNode.Random
		newCurRandom := tableNode[oldCurRandom]
		newCurNode.Random = newCurRandom

		curNode = curNode.Next
	}
	return tableNode[head]
}
