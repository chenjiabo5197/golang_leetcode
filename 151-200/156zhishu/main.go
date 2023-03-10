package main

import "fmt"

func main() {
	fmt.Println(getNum(10))
}

// 统计小于该数字的质数数量
func getNum(a int) int {
	if a <= 2 {
		return 0
	}
	temp := a - 1
	divNum := 2
	for divNum < temp {
		if temp%divNum == 0 {
			return getNum(a - 1)
		} else {
			divNum++
		}
	}
	return 1 + getNum(a-1)
}
