package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 2, 3, 4}
	fmt.Println(triangleNumber(nums))
}

// 611 有效三角形的个数
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	result := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			// 用二分查找找到索引位置
			left := j + 1
			right := len(nums) - 1
			for left <= right {
				mid := (left + right) / 2
				if nums[mid] >= nums[i]+nums[j] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
			//fmt.Printf("left=%d||right=%d\n", left, right)
			// 从循环出来后left=right+1,证明在left=right时，nums[mid] < nums[i] + nums[j]
			// 符合要求的数量从索引j+1到right,包括right
			if nums[right] < nums[i]+nums[j] {
				result += right - j
			}
		}
	}
	return result
}

//func triangleNumber(nums []int) int {
//	sort.Ints(nums)
//	result := 0
//	for i:=0;i<len(nums)-2;i++ {
//		for j:=i+1;j<len(nums)-1;j++ {
//			for k:=j+1;k<len(nums);k++ {
//				if nums[k] >= nums[i] + nums[j] {
//					break
//				}
//				result+=1
//			}
//		}
//	}
//	return result
//}

// 会超时
//func triangleNumber(nums []int) int {
//	result := 0
//	for i:=0;i<len(nums)-2;i++ {
//		for j:=i+1;j<len(nums)-1;j++ {
//			maxLength := nums[i] + nums[j]
//			minLength := int(math.Abs(float64(nums[i]) - float64(nums[j])))
//			for k:=j+1;k<len(nums);k++ {
//				if nums[k] < maxLength && nums[k] > minLength {
//					result += 1
//				}
//			}
//		}
//	}
//	return result
//}
