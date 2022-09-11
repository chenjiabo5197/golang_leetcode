package main

import "fmt"

/*
插入排序
 */

func main() {
	arr := []int{1, 4, 7, 8, 2, 3, 9, 0, 5, 6}
	insertSort(arr)
	fmt.Println(arr)
}

func insertSort(data []int)  {
	for j:=1;j<len(data);j++ {
		temp := data[j]
		for k:=0;k<j;k++ {
			if data[k] > data[j] {  //找到第j个元素要插入的位置
				for l:=j;l>k;l-- {
					if l==0{
						data[l] = 0
					}
					data[l] = data[l-1]  //从k到j所有数字向后挪一位
				}
				data[k] = temp
				break
			}
		}
		fmt.Println(data)
	}
}
