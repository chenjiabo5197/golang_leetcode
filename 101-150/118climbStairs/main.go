package main

import "fmt"

func main() {
	//nums := []int{2, 2, 3, 4}
	fmt.Println(climbStairs(3))
}

//70. 爬楼梯
func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		arr := make([]int, n)
		arr[0] = 1
		arr[1] = 2
		for i := 2; i < n; i++ {
			arr[i] = arr[i-1] + arr[i-2]
		}
		return arr[n-1]
	}
}
