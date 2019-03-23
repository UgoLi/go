// main
package main

import (
	"fmt"
)

func arrDemo1() {
	//1 数组一般声明
	var arr [10]int
	arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr, len(arr))
	// 2 简短声明(并且自动推断长度)
	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr1, len(arr1))
	// 3 下标:值
	arr2 := [...]int{0: 1, 5: 5, 9: 10}
	fmt.Println(arr2, len(arr2))
	// 4 数组遍历
	for i := 0; i < len(arr1); i++ {
		fmt.Print(arr1[i], " ")
	}
	fmt.Println()
	for _, v := range arr1 {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// 5  数组作为函数参数是值传递
	SetValueByArray(arr1)
	fmt.Println(arr1)
	// 6 使用数组指针作为函数的参数，引用传递
	SetValueByArrayPtr(&arr1)
	fmt.Println(arr1)
}
func SetValueByArray(arr [7]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 10
	}
}
func SetValueByArrayPtr(pArr *[7]int) {
	for i := 0; i < len(*pArr); i++ {
		(*pArr)[i] = (*pArr)[i] * 10
	}
}
func main() {
	arrDemo1()
}
