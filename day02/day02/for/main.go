// main
package main

import (
	"fmt"
)

// 求（1~100）的和
func forDemo() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum=", sum)
}

// 省略表达式1
func forDemo2() {
	sum := 0
	i := 1
	for ; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum=", sum)
}

// 省略表达式2
func forDemo3() {
	sum := 0
	i := 1
	for ; ; i++ {
		if i > 100 {
			break
		}
		sum += i
	}
	fmt.Println("sum=", sum)
}

// 省略表达式3
func forDemo4() {
	sum := 0
	i := 1
	for {
		if i > 100 {
			break
		}
		sum += i
		i++
	}
	fmt.Println("sum=", sum)
}
func reverseArray() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("倒置前：", arr)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println("倒置后：", arr)

}
func main() {
	//forDemo()
	//forDemo2()
	//forDemo3()
	//forDemo4()
	reverseArray()
}
