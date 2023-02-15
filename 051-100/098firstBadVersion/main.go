package main

import "fmt"

func main() {
	//nums := []int{2, 2, 3, 4}
	fmt.Println(firstBadVersion(5))
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// 278. 第一个错误的版本
func firstBadVersion(n int) int {
	left := 0
	right := n
	for left < right {
		mid := left + (right-left)/2
		curVersion := isBadVersion(mid)
		if curVersion { //为true表示当前版本大于等于错误版本，所以应该向前继续找,不减1因为有可能当前状态是第一个坏版本
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

//超时
//func firstBadVersion(n int) int {
//	curVersion := isBadVersion(n)
//	for curVersion {
//		curVersion = isBadVersion(n-1)
//		if !curVersion {
//			return n
//		}else {
//			n -= 1
//		}
//	}
//	return n
//}

func isBadVersion(version int) bool {
	bad := 4
	if version >= bad {
		return true
	}
	return false
}
