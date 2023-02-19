package main

import "fmt"

func main() {
	nums := []int{2, 3, 2}
	fmt.Println(rob(nums))
}

//213. 打家劫舍 II
/*
	一维数组arr，arr[i]表示到i为止，可以获取大最大金额
	arr[i] = max(arr[i-2]+nums[i], arr[i-1])
	房屋收尾相连，则偷了第一间不能偷最后一间，范围为0~len(bums)-2
				  相偷最后一间不能偷第一间，范围为1~len(nums)-1
*/
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	arr1 := make([]int, len(nums)-1) // 偷第一间
	arr2 := make([]int, len(nums)-1) // 偷最后一间
	arr1[0] = nums[0]
	arr2[0] = nums[1]
	for i := 1; i < len(arr1); i++ {
		if i == 1 {
			arr1[i] = max(nums[0], nums[1])
			arr2[i] = max(nums[1], nums[2]) // 到此处nums数组长度最小为3，所以不会越界
		} else if i >= 2 {
			arr1[i] = max(arr1[i-2]+nums[i], arr1[i-1])
			arr2[i] = max(arr2[i-2]+nums[i+1], arr2[i-1])
		}
	}
	return max(arr1[len(arr1)-1], arr2[len(arr2)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
