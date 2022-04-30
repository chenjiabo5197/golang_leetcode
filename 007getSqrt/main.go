package main

import (
	"fmt"
	"strconv"
)

/*
计算整数的算术平方根
 */

func main() {
	n := 4
	fmt.Println(mySqrt(n))
}

func mySqrt(x int) int {
	result := 0
	strX := strconv.Itoa(x)
	resultStr := ""
	if len(strX) % 2 != 0 {
		strX = "0"+strX
	}
	reminderNum := 0
	for i:=0;i<=len(strX)-2;i+=2 {
		partStr := strX[i:i+2]
		partInt,_ := strconv.Atoi(partStr)
		nowNum := 0
		partInt += reminderNum*100
		for j:=0;j<10;j++ {
			find := false
			if i != 0 {
				leftNum,_ := strconv.Atoi(resultStr)
				nowNum = leftNum*2*10+j
				if nowNum * j <= partInt && (nowNum+1)*(j+1) >partInt {
					find = true
				}
			}else {
				if j *j <= partInt && (j+1)*(j+1)>partInt {
					find = true
				}
			}
			if find {
				resultStr += strconv.Itoa(j)
				if i == 0 {
					reminderNum = partInt-(j*j)
				}else {
					reminderNum = partInt-(nowNum*j)
				}
				break
			}
		}
	}
	result,_ = strconv.Atoi(resultStr)
	return result
}

