package main

/*
最小栈
 */

func main() {
	
}

type MinStack struct {
	data []int
	topPoint int
	minPoint int
}

func Constructor() MinStack {
	m1 := MinStack{
		data:     make([]int, 0),
		topPoint: -1,
		minPoint: -1,
	}
	return m1
}


func (this *MinStack) Push(val int)  {
	this.data = append(this.data, val)
	this.topPoint ++
}


func (this *MinStack) Pop()  {
	if this.topPoint == -1 {
		return
	}
	this.data = this.data[:len(this.data)-1]
	this.topPoint--
}


func (this *MinStack) Top() int {
	if this.topPoint == -1 {
		return 0
	}
	return this.data[this.topPoint]
}


func (this *MinStack) GetMin() int {
	min := 0
	for i:=0;i<len(this.data);i++ {
		if i == 0 {
			min = this.data[0]
		}else if min > this.data[i]{
			min = this.data[i]
		}
	}
	return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
