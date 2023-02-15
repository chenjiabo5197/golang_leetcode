package main

import "fmt"

func main() {
	//nums := []int{2, 2, 3, 4}
	fmt.Println(isPerfectSquare(16))
}

//367. 有效的完全平方数
func isPerfectSquare(num int) bool {
	min := 0
	max := num
	for min <= max {
		mid := min + (max-min)/2
		if mid*mid == num {
			return true
		} else if mid*mid > num {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return false
}
