package main

import (
	"fmt"
)

/*
计算阶乘末尾0的个数，即计算阶乘因数中5的个数
 */

func main() {
	n := 30
	fmt.Println(trailingZeroes(n))
}

func trailingZeroes(n int) int {
	result := 0
	for{
		if n < 5 {
			break
		}
		temp := n
		for{
			if temp < 5 {
				break
			}
			if temp % 5 == 0 {
				 result++
				temp = temp / 5
			}else {
				break
			}
		}
		n --
	}
	return result
}


