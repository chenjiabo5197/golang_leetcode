package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	data := "[1,2,3,null,null,4,5,6,7]"
	//ser := Constructor()
	deser := Constructor()
	root := deser.deserialize(data)
	data1 := deser.serialize(root)
	fmt.Println(data1)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

// 297 二叉树的序列化与反序列化
type Codec struct {

}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	storeyNode := make([]*TreeNode, 0)
	storeyNode = append(storeyNode, root)
	result := "["
	for {
		result += node2String(storeyNode)
		storeyNode = getNextStoreyNode(storeyNode)
		if !hasNextStoreyNode(storeyNode) {
			break
		}
	}
	result = result[:len(result)-1]   // 去除最后一个,
	result += "]"
	return result
}

func node2String(data []*TreeNode) string {
	result := ""
	for index := range data {
		if data[index] == nil {
			result += "null"
		}else {
			result += strconv.Itoa(data[index].Val)
		}
		result += ","
	}
	return result
}

func getNextStoreyNode(nodes []*TreeNode) []*TreeNode {
	nextStoreyNodes := make([]*TreeNode, 0)
	for index := range nodes {
		if nodes[index] == nil {
			continue
		}
		nextStoreyNodes =append(nextStoreyNodes, nodes[index].Left)
		nextStoreyNodes =append(nextStoreyNodes, nodes[index].Right)
	}
	return nextStoreyNodes
}

// 判断现在队列中的节点是否有下一层节点
func hasNextStoreyNode(nodes []*TreeNode) bool {
	for index := range nodes {
		if nodes[index] == nil {
			continue
		}
		if nodes[index] != nil {
			return true
		}
	}
	return false
}


// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 2 {
		return nil
	}
	valueSlice := string2Slice(data)
	nodeAddr := make([]*TreeNode, 1)
	root := &TreeNode{
		Val:   valueSlice[0],
	}
	nodeAddr[0] = root
	nodeAddrIndex := 0   //指示当前的父节点的索引
	preNodeRight, preNodeLeft := false, false   // 指示父节点的左节点是否赋值
	for index:=1;index<len(valueSlice);index++ {
		var node *TreeNode
		if valueSlice[index] == -1001 {   // 该节点为空
			node = nil
		}else {
			node = &TreeNode{
				Val: valueSlice[index],
			}
		}
		nodeAddr = append(nodeAddr, node)
		for nodeAddr[nodeAddrIndex]==nil {
			nodeAddrIndex++
		}
		if !preNodeLeft{
			nodeAddr[nodeAddrIndex].Left = node
			preNodeLeft = true
			continue
		}
		if !preNodeRight {
			nodeAddr[nodeAddrIndex].Right = node
			preNodeRight = true
		}
		if preNodeLeft && preNodeRight {
			nodeAddrIndex++
			preNodeRight, preNodeLeft = false, false
		}

	}
	return root
}

func string2Slice(data string) []int {
	data = data[1:len(data)-1]  //去除两边括号
	valueStr := strings.Split(data, ",")
	valueInt := make([]int, 0)
	for index := range valueStr {
		value, err := strconv.Atoi(valueStr[index])
		if err != nil {
			valueInt = append(valueInt, -1001)
		}else {
			valueInt = append(valueInt, value)
		}
	}
	return valueInt
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */