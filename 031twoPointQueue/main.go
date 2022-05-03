package main

import "fmt"

/*
双端队列
 */

func main() {
	q := Constructor(4)
	q.InsertLast(1)
	q.InsertLast(2)
	q.InsertFront(3)
	q.InsertFront(4)
	fmt.Println(q.GetRear())
}

type MyCircularDeque struct {
	data []int
	maxLength int
	headPoint int
}


func Constructor(k int) MyCircularDeque {
	queue := MyCircularDeque{
		data:      make([]int, 0),
		maxLength:  k,
		headPoint: -1,
	}
	return queue
}


func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data = append(this.data, value)
	this.headPoint ++
	return true
}


func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	tempSlice := []int{value}
	this.data = append(tempSlice, this.data...)
	this.headPoint ++
	return true
}


func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.data = this.data[:len(this.data) - 1]
	this.headPoint --
	return true
}


func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.data = this.data[1:]
	this.headPoint --
	return true
}


func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.headPoint]
}


func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[0]
}


func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.data) == -1
}


func (this *MyCircularDeque) IsFull() bool {
	return len(this.data) == this.maxLength
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
