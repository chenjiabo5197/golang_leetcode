package main

import "fmt"

func main() {
	q := Constructor()
	q.Push(1)
	fmt.Println(q.Peek())
}

type MyQueue struct {
	pushStack *Mystack
	popStack *Mystack
}


func Constructor() MyQueue {
	q := MyQueue{
		pushStack: &Mystack{
			data:      make([]int, 0),
			headPoint: -1,
		},
		popStack:  &Mystack{
			data:      make([]int, 0),
			headPoint: -1,
		},
	}
	return q
}


func (this *MyQueue) Push(x int)  {
	this.pushStack.push(x)
}


func (this *MyQueue) Pop() int {
	if this.popStack.isEmpty() {
		for {
			if this.pushStack.isEmpty() {
				break
			}
			temp := this.pushStack.peak()
			this.pushStack.pop()
			this.popStack.push(temp)
		}
	}
	if this.Empty() {
		return 0
	}
	temp := this.popStack.peak()
	this.popStack.pop()
	return temp
}


func (this *MyQueue) Peek() int {
	if this.popStack.isEmpty() {
		for {
			if this.pushStack.isEmpty() {
				break
			}
			temp := this.pushStack.peak()
			this.pushStack.pop()
			this.popStack.push(temp)
		}
	}
	if this.Empty() {
		return 0
	}
	return this.popStack.peak()
}


func (this *MyQueue) Empty() bool {
	return this.popStack.isEmpty() && this.pushStack.isEmpty()
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

type Mystack struct {
	data []int
	headPoint int
}

func (stack *Mystack) push (value int) {
	stack.data = append(stack.data, value)
	stack.headPoint ++
}

func (stack *Mystack) pop () {
	if stack.isEmpty() {
		return
	}
	stack.data = stack.data[:len(stack.data)-1]
	stack.headPoint--
}
func (stack *Mystack) isEmpty()  bool {
	return len(stack.data) == 0
}

func (stack *Mystack) peak() int {
	if stack.isEmpty() {
		return 0
	}
	return stack.data[stack.headPoint]
}
