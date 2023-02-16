package main

import "fmt"

func main() {
	//nums := []int{2, 2, 3, 4}
	fmt.Println(hammingWeight(00000000000000000000000000001011))
}

//191. 位1的个数
func hammingWeight(num uint32) int {
	result := 0
	for num != 0 {
		// 会将最右侧为1的消除
		num = num & (num - 1)
		result += 1
	}
	return result
}
