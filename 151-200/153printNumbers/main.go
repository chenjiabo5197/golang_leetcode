package main

import "fmt"

func main() {
	n := 2
	fmt.Println(printNumbers(n))
}

//剑指 Offer 17. 打印从1到最大的n位数
func printNumbers(n int) []int {
	result := make([]int, 0)
	bitNumber := 1
	start := 1
	for bitNumber <= n {
		result = append(result, start)
		start++
		bitNumber = getBit(start)
	}
	return result
}

func getBit(number int) int {
	bitNumber := 1
	for number/10 != 0 {
		bitNumber++
		number = number / 10
	}
	return bitNumber
}
