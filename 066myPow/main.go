package main

import "fmt"

func main() {
	x := 2.00000
	n := 10
	fmt.Println(myPow(x, n))
}


// 数值的整数次方
func myPow(x float64, n int) float64 {
	if n == 0{
		return 1
	}else if n < 0 {
		return 1.0/recursion(x, n)
	}
	return recursion(x, n)
}

func recursion (x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	temp := recursion(x, n/2)
	if n%2 == 0{
		return temp*temp
	}else {
		return temp*temp*x
	}
}