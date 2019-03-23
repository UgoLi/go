// main
package main

import (
	"fmt"
)

// 一般函数，求两个数的最大值
func Max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

//  1 一般函数，求两个数的最小值
func Min(a, b int) (min int) {
	if a <= b {
		min = a
	} else {
		min = b
	}
	return
}

// 2 多返回值函数
func AddAndSub(a, b int) (int, int) {
	return a + b, a - b
}
func SubAndAdd(a, b int) (sub, add int) {
	sub = a - b
	add = a + b
	return
}

// 3 可变参数
func Sum(nums ...int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println("sum=", sum)
}

// 4 递归函数(20分钟)
// 4.1 求阶乘，例如:5!=5*4*3*2*1=120, f(n)=n*f(n-1)
func jiechen(num int) int {
	if num == 1 {
		return 1
	}
	return num * jiechen(num-1)
}

// 4.2 求斐波那契数列，例如: 1 1 2 3 5 8 13 21 34 55...
// f(n)=f(n-1)+f(n-2)
func feibo(num int) int {
	if num <= 2 {
		return 1
	}
	return feibo(num-1) + feibo(num-2)
}
func main() {
	for i := 1; i <= 20; i++ {
		fmt.Print(feibo(i), " ")
	}
	fmt.Println()
	/*
		fmt.Println(Max(10, 5))
		fmt.Println(Min(-10, -5))
		fmt.Println(AddAndSub(10, 5))
		fmt.Println(SubAndAdd(-10, -5))
		Sum(1)
		Sum(1, 2, 3)
		Sum(1, 2, 3, 4, 5)
		slice := []int{1, 2, 3, 4, 5, 6, 7}
		Sum(slice...)
		fmt.Println(jiechen(5))
	*/
}
