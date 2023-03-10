package main

import "fmt"

func main() {
	x := 2.00000
	n := 10
	fmt.Println(myPow(x, n))
}

//剑指 Offer 16. 数值的整数次方
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n%2 == 0 {
		temp := myPow(x, n/2)
		return temp * temp
	} else {
		temp := myPow(x, n/2)
		return temp * temp * x
	}
}
