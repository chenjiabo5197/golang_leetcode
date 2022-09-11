package main

import "fmt"

/*
队列的基础信息
 */

func main() {
	q := MyQueue{
		data: make([]int, 0),
		head: -1,
		rear: -1,
	}
	q.appendQueue(1)
	q.appendQueue(2)
	q.appendQueue(3)
	fmt.Println(q.getHead())
	fmt.Println(q.size())
	q.remove()
	fmt.Println(q.getHead())
	fmt.Println(q.size())
	q.remove()
	fmt.Println(q.getHead())
	fmt.Println(q.size())
	q.remove()
	fmt.Println(q.getHead())
	fmt.Println(q.size())
}

type MyQueue struct {
	data []int
	head int
	rear int
}

// 向队列中添加元素
func (q *MyQueue) appendQueue (value int) {
	if q.isEmpty() {
		q.rear ++
	}
	// 向尾部增加元素，从头部出去，所以将新增的元素放在索引为0的位置
	tempSlice := make([]int, 0)
	tempSlice = append(tempSlice, value)
	q.data = append(tempSlice, q.data...)
	q.head ++
}

// 获取队首元素
func (q *MyQueue) getHead() int {
	if q.isEmpty() {
		return 0
	}
	return q.data[q.head]
}

// 删除队首元素
func (q *MyQueue) remove() {
	if q.isEmpty() {
		return
	}
	q.data = q.data[:len(q.data)-1]
	q.head--
}

// 判断队列是否为空
func (q *MyQueue) isEmpty() bool {
	return len(q.data) == 0
}

// 获取队列中元素个数
func (q *MyQueue) size() int {
	return len(q.data)
}
