package main

import (
	"errors"
	"fmt"
)

/*
每日温度
 */

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
}

//func dailyTemperatures(temperatures []int) []int {
//	result := make([]int, 0)
//	for i:=0;i<len(temperatures);i++ {
//		index := 0
//		flag := false
//		for j:=i+1;j<len(temperatures);j++ {
//			if temperatures[i] >= temperatures[j] {
//				index++
//			}else {
//				flag = true
//				index++
//				break
//			}
//		}
//		if flag {
//			result = append(result, index)
//		}else {
//			result = append(result, 0)
//		}
//	}
//	return result
//}

//用栈来做，只用遍历一次切片即可，栈中保存节点的下标值
func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := &MyStack{
		data:      make([]int, 0),
		stackTop:  -1,
		stackFoot: -1,
	}
	for i:=0;i<len(temperatures);i++ {
		for {
			if stack.isEmpty() || temperatures[stack.peak()] >= temperatures[i] {
				break
			}
			result[stack.peak()] = i - stack.peak()
			stack.pop()
		}
		stack.push(i)
	}
	return result
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

