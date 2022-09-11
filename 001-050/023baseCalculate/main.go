package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	s := "(3)+1"
	//fmt.Println(int(s[1])-48)
	fmt.Println(calculate(s))
}

//func calculate(s string) int {
//	stackNum := &MyStack{
//		data:      make([]string, 0),
//		stackTop:  -1,
//		stackFoot: -1,
//	}
//	stackOper := &MyStack{
//		data:      make([]string, 0),
//		stackTop:  -1,
//		stackFoot: -1,
//	}
//	for i:=0;i<len(s);i++ {
//		data := string(s[i])
//		if i == 0 && data == "-" {
//			stackNum.push("0")
//			stackOper.push("-")
//		}else if data == " "{
//			continue
//		}else if data == "+" || data == "-" {
//			topString, _ := stackNum.peak()
//			if topString == "(" {
//				stackNum.push("0")
//				stackOper.push("-")
//			}
//			stackOper.push(data)
//		}else if data == ")" {
//			for {
//				operator, errOperat := stackOper.pop()
//				num2, _ := stackNum.pop()
//				num1, _ := stackNum.pop()
//				if _, err := strconv.Atoi(num1); err != nil && errOperat != nil {
//					stackNum.push(num2)
//					stackOper.push(operator)
//					break
//				}
//				tempResult := doOperator(num1, num2, operator)
//				topString, _ := stackNum.peak()
//				if topString == "(" {
//					stackNum.pop()
//					stackNum.push(strconv.Itoa(tempResult))
//					break
//				}
//				stackNum.push(strconv.Itoa(tempResult))
//			}
//		}else if data == "(" {
//			stackNum.push(data)
//		}else {
//			data, i = getAllNum(data, s, i)
//			stackNum.push(data)
//		}
// 	}
//	if stackOper.isEmpty() {
//		topString, _ := stackNum.peak()
//		resultInt, _ := strconv.Atoi(topString)
//		return resultInt
//	}
// 	//将数据栈和运算符栈弹出，重新转入两个新的栈中，从左往右计算
// 	newNumStack := &MyStack{
//		data:      make([]string, 0),
//		stackTop:  -1,
//		stackFoot: -1,
//	}
//	newOperaStack := &MyStack{
//		data:      make([]string, 0),
//		stackTop:  -1,
//		stackFoot: -1,
//	}
// 	for {
// 		if stackNum.isEmpty() && stackOper.isEmpty() {
// 			break
//		}
//		if numStr, err := stackNum.pop(); err == nil {
//			newNumStack.push(numStr)
//		}
//		if operaStr, err := stackOper.pop(); err == nil {
//			newOperaStack.push(operaStr)
//		}
//	}
//	for {
//		if newNumStack.size() == 1 {
//			topString, _ := newNumStack.peak()
//			resultInt, _ := strconv.Atoi(topString)
//			return resultInt
//		}
//		operator, _ := newOperaStack.pop()
//		num1, _ := newNumStack.pop()
//		num2, _ := newNumStack.pop()
//		tempResult := doOperator(num1, num2, operator)
//		newNumStack.push(strconv.Itoa(tempResult))
//	}
//}

type MyStack struct {
	data []string
	stackTop int
	stackFoot int
}

/*
将结果暂时储存在栈中，栈顶为左括号外面的符号   结果为res*sign  + 可表sign为1,- 可表sign为-1
 */
func calculate(s string) int {
	stack := &MyStack{
		data:      make([]string, 0),
		stackTop:  -1,
		stackFoot: -1,
	}
	res := 0  //暂存结果
	sign := 1 //区分正负号
	for i:=0;i<len(s);i++ {
		if s[i] >= 48 && s[i] <= 57 {
			data, _ :=  strconv.Atoi(string(s[i]))
			for {
				if i+1 >= len(s) || s[i+1] < 48 || s[i+1] > 57{
					break
				}
				data = data*10 + int(s[i+1])-48
				i++
			}
			res = res + sign * data
		}else if string(s[i]) == "-" {
			sign = -1
		}else if string(s[i]) == "+" {
			sign = 1
		}else if string(s[i]) == "(" {
			stack.push(strconv.Itoa(res))
			stack.push(strconv.Itoa(sign))
			res = 0    //压栈之后重新初始化
			sign = 1
		}else if string(s[i]) == ")" {
			//从栈取出元素，与现在的值进行计算
			signStr, _ := stack.peak()
			stack.pop()
			sign, _ = strconv.Atoi(signStr)
			preResStr, _ := stack.peak()
			stack.pop()
			preRes, _ := strconv.Atoi(preResStr)
			res = res * sign + preRes
		}
	}
	return res
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
