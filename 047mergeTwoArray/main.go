package main

import "fmt"

/*
合并两个有序数组
 */

func main() {
	nums1 := []int{4, 5, 6, 0, 0, 0}
	m := 3
	nums2 := []int{1, 2, 3}
	n := 3
	merge(nums1, m, nums2, n)
}

//从前向后遍历
//func merge(nums1 []int, m int, nums2 []int, n int)  {
//	if m == 0 && n != 0 {
//		for i:=0;i<n;i++ {
//			nums1[i] = nums2[i]
//		}
//	}else if m !=0 && n != 0 {
//		index1 := 0
//		index2 := 0
//		for i:=0;i<len(nums1);i++ {
//			if index1 < m && index2 < n {
//				if nums1[index1] > nums2[index2] {
//					//将nums1中从第i个数字到第m个数字都向后挪一位
//					for j:=m;j>i;j-- {
//						nums1[j] = nums1[j-1]
//					}
//					index1++  //index1加1表示需要从nums1中本来的数字开始遍历，如果不加，会从加入的数字开始比较
//					m++  //挪完之后m加1，表示nums1中数字遍历需要多往后遍历一位
//					nums1[i] = nums2[index2]
//					index2++
//					continue
//				}
//				index1++
//			}else if index1 >= m {  //此处只用考虑数组1排序完（  ）之后的事，数组2排序完即结束
//				nums1[i] = nums2[index2]
//				index2++
//			}
//		}
//	}
//	fmt.Println(nums1)
//}

//从nums1的数组最后一位开始比较，则只需要遍历一次
func merge(nums1 []int, m int, nums2 []int, n int) {
	index1 := m-1
	index2 := n-1
	//curIndex := m+n-1
	for i:=len(nums1)-1;i>=0;i-- {
		if index2 < 0 {  //nums2中数组排序完则整个排序结束
			break
		}
		if index1 < 0 || nums1[index1] < nums2[index2]{
			nums1[i] = nums2[index2]
			index2--
		}else {
			nums1[i] = nums1[index1]
			index1--
		}
	}
	fmt.Println(nums1)
}