package main

import (
	"fmt"
	"strconv"
)

/*
移掉k位数字
 */

func main() {
	num := "1432219"
	k := 3
	fmt.Println(removeKdigits(num, k))
}

// 目前不能成功处理要移除的数字在末尾的情况
//func removeKdigits(num string, k int) string {
//	if len(num) <= k {
//		return "0"
//	}else {
//		result := ""
//		lastIndex := -1  //上一轮循环找到的数的索引，初始为-1
//		for k > 0 {
//			index := lastIndex+1
//			for i:=lastIndex+2;i<=lastIndex+k+1 && i<len(num);i++ {
//				if num[i] < num[index] {
//					index = i
//				}
//			}
//			//将此轮前面应该删除的数字删除
//			if lastIndex == -1 {  //如果是第一次删除数字，因为初始化为-1
//				k = k - index
//
//			}else {
//				k = k - (index - lastIndex - 1)
//			}
//			// 更新lastIndex
//			lastIndex = index
//			result += string(num[index])
//			if len(result) == len(num) - k {
//				break
//			}
//			fmt.Println(index, result,k)
//		}
//		result += num[lastIndex+1:]
//		for k,v := range []rune(result) {  //将前面的0去掉
//			//fmt.Println(v)
//			if v != 48 {
//				break
//			}
//			if k >= len(result) -1 {
//				return "0"
//			}
//			result = result[k+1:]
//		}
//		if len(result) <= 0 {
//			return "0"
//		}
//		return result
//	}
//}

// 用栈来做，每次都判断该位置数字是否应该移除
func removeKdigits(num string, k int) string {
	numStack := []int{0}  // 用切片模拟栈,初始放入一个0，防止切片越界
	for i:=0;i<len(num); {
		if numStack[len(numStack)-1] <= int(num[i]-48) || k <=0 {  //数字被移除完,剩下的所有数字都入栈
			numStack = append(numStack, int(num[i]-48))
			i++
		}else {
			numStack = numStack[:len(numStack)-1]  //从栈中取出一个数字
			k--
		}
	}
	for k > 0 {
		numStack = numStack[:len(numStack)-1]  //从栈中取出一个数字
		k--
	}
	result := ""
	flag := true  //标志数据开头
	for i:=1;i<len(numStack);i++ {   //因为前面填充了一个0
		if numStack[i] != 0 {
			flag = false
		}
		if flag && numStack[i] == 0 {
			continue
		}
		result = result + strconv.Itoa(numStack[i])
	}
	if len(result) <= 0 {
		return  "0"
	}
	return result
}