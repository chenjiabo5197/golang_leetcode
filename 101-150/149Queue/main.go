package main

//剑指 Offer 09. 用两个栈实现队列
type CQueue struct {
	Data []int
	Head int
	Tail int
}

func Constructor() CQueue {
	c := CQueue{}
	c.Data = make([]int, 0)
	c.Head = 0
	c.Tail = -1
	return c
}

func (this *CQueue) AppendTail(value int) {
	this.Data = append(this.Data, value)
	this.Tail++
}

func (this *CQueue) DeleteHead() int {
	if this.Tail == -1 || this.Head >= len(this.Data) {
		return -1
	}
	temp := this.Data[this.Head]
	this.Head++
	return temp
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
