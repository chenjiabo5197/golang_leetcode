package main

import "fmt"

func main() {
	//nums := []int{2, 2, 3, 4}
	fmt.Println(isPowerOfTwo(16))
}

//231. 2 的幂
func isPowerOfTwo(n int) bool {
	return (n > 0) && (n&(n-1) == 0)
}

//func isPowerOfTwo(n int) bool {
//	if n <= 0 {
//		return false
//	}
//	if n == 1 {
//		return true
//	}
//	for n >= 2 {
//		if n % 2 != 0 {
//			return false
//		}else {
//			n = n / 2
//		}
//	}
//	if n == 1 {
//		return true
//	}
//	return false
//}
