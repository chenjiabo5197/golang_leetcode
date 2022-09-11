package main

import (
	"fmt"
	"strconv"
)

/*
根据给定的字符串计算最后得分
 */

func main() {
	str := []string{"5","2","C","D","+"}
	fmt.Println(calPoints(str))
}

func calPoints(ops []string) int {
	result := make([]int, 0)
	for i:=0;i<len(ops);i++ {
		if ops[i] == "C" {
			result = result[:len(result)-1]
			continue
		}
		if ops[i] == "D" {
			temp := 2 * result[len(result)-1]
			result = append(result, temp)
			continue
		}
		if ops[i] == "+" {
			temp1 := result[len(result)-1]
			temp2 := result[len(result)-2]
			result = append(result, temp1+temp2)
			continue
		}
		temp,_ := strconv.Atoi(ops[i])
		result = append(result, temp)
	}
	resp := 0
	for i:=0;i<len(result);i++{
		resp += result[i]
	}
	return resp
}
