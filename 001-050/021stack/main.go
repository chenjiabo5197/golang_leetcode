package main

import (
	"errors"
	"fmt"
)

/*
栈的基础
 */

func main() {
	stack := &MyStack{
		data:      make([]string,0),
		stackTop:  -1,
		stackFoot: -1,
	}
	stack.push("1")
	stack.push("2")
	stack.push("3")
	stack.push("4")
	fmt.Println(stack.data)
	fmt.Println(stack.stackTop)
	fmt.Println(stack.peak())
	fmt.Println(stack.size())

	stack.pop()
	stack.pop()
	stack.pop()
	stack.pop()
	stack.pop()
	fmt.Println(stack.peak())
	fmt.Println(stack.size())
}

type MyStack struct {
	data []string
	stackTop int
	stackFoot int
}

//向栈里添加元素
func (stack *MyStack) push (s string) {
	stack.data = append(stack.data, s)
	stack.stackTop ++
}

//获取栈顶元素
func (stack *MyStack) peak() (string, error) {
	if stack.isEmpty() {
		err := errors.New("栈为空")
		return "", err
	}
	s := stack.data[stack.stackTop]
	//stack.data = stack.data[:len(stack.data)-1]
	return s, nil
}

//弹出栈顶元素
func (stack *MyStack) pop() (string,error) {
	if stack.isEmpty() {
		err := errors.New("栈为空")
		return "", err
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
