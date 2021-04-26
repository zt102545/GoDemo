package main

import "fmt"

//冒泡排序
func BubbleSort(list []int) []int {
	for i:=0;i< len(list);i++{
		for j:=i;j<len(list);j++{
			if list[i]>list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}

func main() {
	list := [...]int{1, 2, 3, 4, 5, 2, 5}
	var newList = BubbleSort(list[:])
	fmt.Printf("%v", newList)
}
