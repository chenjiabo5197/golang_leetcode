package main

import (
	"errors"
	"fmt"
)

/*
判断输入的括号是否有效
 */

func main() {
	s := "([)]"
	fmt.Println(isValid(s))
}

type MyStack struct {
	data []string
	stackTop int
	stackFoot int
}

func isValid(s string) bool {
	stack := &MyStack{
		data:      make([]string, 0),
		stackTop:  -1,
		stackFoot: -1,
	}
	for _, value := range []rune(s) {
		data := string(value)
		if data == "(" {
			stack.push(data)
		}else if data == "[" {
			stack.push(data)
		}else if data == "{" {
			stack.push(data)
		}else {
			topString, err := stack.peak()
			if err != nil {
				return false
			}
			if (data == ")" && topString == "(") || (data == "}" && topString == "{") || (data == "]" && topString == "[" ) {
				stack.pop()
			}else {
				return false
			}
		}

	}
	return stack.isEmpty()
}

//向栈里添加元素
func (stack *MyStack) push (s string) {
	stack.data = append(stack.data, s)
	stack.stackTop ++
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

func (stack *MyStack) peak() (string, error) {
	if stack.isEmpty() {
		err := errors.New("栈为空")
		return "", err
	}
	s := stack.data[stack.stackTop]
	//stack.data = stack.data[:len(stack.data)-1]
	return s, nil
}

func (stack *MyStack) isEmpty() bool{
	return stack.stackTop == -1
}
