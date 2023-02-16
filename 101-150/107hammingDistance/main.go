package main

import "fmt"

func main() {
	x := 1
	y := 4
	fmt.Println(hammingDistance(x, y))
}

//461. 汉明距离
func hammingDistance(x int, y int) int {
	andValue := x ^ y
	result := 0
	for andValue > 0 {
		tempValue := andValue - 1
		andValue = andValue & tempValue
		result += 1
	}
	return result
}
