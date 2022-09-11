package main

import (
	"errors"
	"fmt"
)

/*
验证栈序列
 */

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	fmt.Println(validateStackSequences(pushed, popped))
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := &MyStack{
		data:      make([]int, 0),
		stackTop:  -1,
		stackFoot: -1,
	}
	index := 0
	for i:=0;i<len(pushed);i++ {
		stack.push(pushed[i])
		for {
			if stack.isEmpty() || stack.peak() != popped[index] {
				break
			}
			stack.pop()
			index++
		}
	}
	if stack.isEmpty() {
		return true
	}
	return false
}

type MyStack struct {
	data []int
	stackTop int
	stackFoot int
}

//向栈里添加元素
func (stack *MyStack) push (s int) {
	stack.data = append(stack.data, s)
	stack.stackTop ++
}

//获取栈顶元素
func (stack *MyStack) peak() int {
	if stack.isEmpty() {
		return 0
	}
	s := stack.data[stack.stackTop]
	//stack.data = stack.data[:len(stack.data)-1]
	return s
}

//弹出栈顶元素
func (stack *MyStack) pop() (int,error) {
	if stack.isEmpty() {
		err := errors.New("栈为空")
		return 0, err
	}
	s := stack.data[stack.stackTop]
	stack.data = stack.data[:len(stack.data)-1]
	stack.stackTop --
	return s, nil
}

//判断栈是否为空
func (stack *MyStack) isEmpty() bool{
	return stack.stackTop == -1
}

//获取栈中元素个数
func (stack *MyStack) size() int {
	return stack.stackTop + 1
}

