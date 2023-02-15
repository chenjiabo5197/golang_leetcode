package main

import "fmt"

/*
柠檬水找零
 */

func main() {
	data := []int{5,5,10,10,20}
	billTable := make(map[int]int)
	v, ok := billTable[1]
	fmt.Println(v, ok)
	billTable[1] = v+1
	v, ok = billTable[1]
	fmt.Println(v, ok)
	fmt.Println(lemonadeChange(data))
}

func lemonadeChange(bills []int) bool {
	billTable := make(map[int]int)
	for i:=0;i<len(bills);i++ {
		if bills[i] == 5 {
			value, _ := billTable[5]
			billTable[5] = value+1
		}else if bills[i] == 10 {
			value5, ok5 := billTable[5]
			value10, _ := billTable[10]
			if !ok5 || value5 <= 0{
				return false
			}else {
				billTable[10] = value10+1
				billTable[5] = value5-1
			}
		}else {
			value5, _ := billTable[5]
			value10, _ := billTable[10]
			if value5 <=0 || (value10 <= 0 && value5 <=2) {
				return false
			}else {
				if value10 > 0 && value5 > 0 {
					billTable[10] = value10-1
					billTable[5] = value5-1
				}else {
					billTable[5] = value5-3
				}
			}
		}
	}
	value5, _ := billTable[5]
	value10, _ := billTable[10]
	if value5 >= 0 && value10 >= 0 {
		return true
	}else {
		return false
	}
}