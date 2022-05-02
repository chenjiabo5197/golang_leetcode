package main

import (
	"errors"
	"fmt"
	"math"
)

/*
获得能截取雨水的总量
 */

func main() {
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(height))

}

/*
1、建立一个递增的栈
2、判断要入栈的元素与栈顶元素的大小
	若小于等于栈顶元素，则直接入栈
	若大于栈顶元素，则先弹出当前栈顶元素，用弹出的元素做底，高度是当前元素与弹出后的栈顶元素的最小值得到的，然后再减去底的高度
														宽度是当前元素序号减去栈顶元素的序号得到的
 */
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	stack := &MyStack{   //栈中保存元素在切片中的下标，元素值可以通过下标O(1)时间内从切片中获取
		data:      make([]int, 0),
		stackTop:  -1,
		stackFoot: -1,
	}
	res := 0
	for i:=0;i<len(height);i++ {
		if stack.isEmpty() || height[i] <= height[stack.peak()] {
			stack.push(i)
		}else {
			for {
				if stack.isEmpty() || height[i] <= height[stack.peak()] {
					break
				}
				bottom := stack.peak()   //获取矩形底部的高度
				stack.pop()
				if !stack.isEmpty() {
					h := int(math.Min(float64(height[stack.peak()]), float64(height[i]))) - height[bottom]
					w := i - stack.peak() - 1
					res += h * w
				}
			}
			stack.push(i)
		}
	}
	return res
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
