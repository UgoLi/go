// main
package main

import (
	"fmt"
)

func sliceDemo() {
	// 1 声明切片
	var slice1 []int
	fmt.Println(slice1)
	// 2 基于数组创建切片
	var arr1 [7]int = [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("数组", arr1)
	slice1 = arr1[:]
	fmt.Println("切片", slice1)
	// 3 基于切片创建切片
	slice2 := slice1[1:5]
	fmt.Println("切片2", slice2)
	slice2[1] = 30
	fmt.Println(arr1)
	fmt.Println(slice1)
	fmt.Println(slice2)
	// 4 直接使用内置的make函数创建切片
	slice3 := make([]int, 5)
	fmt.Println(slice3, len(slice3), cap(slice3))
	// 使用时，下标不能越界，最大值是len()-1
	slice3[len(slice3)-1] = 5
	fmt.Println(slice3)
	//slice3[len(slice3)] = 5//赋值不能越界
	//追加元素时，尽量不要超出cap，超出时，会有内存操作。
	// 5 空切片，长度为零
	var slice4 []int
	fmt.Println(len(slice4) == 0)
	// 6 切片不能比较
	//fmt.Println(slice1 == slice4)//error
	// 7 追加元素
	fmt.Println(slice1)
	slice1 = append(slice1, 8)
	fmt.Println(slice1)
	// 8 追加slice
	slice5 := []int{9, 10, 11}
	//将slice5打散，不是[9,10,11]，而是9,10,11
	slice1 = append(slice1, slice5...)
	fmt.Println(slice1)
	// 9 切片复制
	slice6 := []int{1, 2, 3, 4, 5}
	slice7 := []int{10, 20, 30}
	//copy(slice6, slice7)
	copy(slice7, slice6)
	fmt.Println(slice6, slice7)

}
func SetValueBySlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = slice[i] * 10
	}
}
func main() {
	//sliceDemo()
	//10 将slice作为函数的参数是引用传递
	arr := [...]int{1, 2, 3, 4, 5}
	SetValueBySlice(arr[:])
	fmt.Println(arr)
}
