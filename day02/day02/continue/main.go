// main
package main

import (
	"fmt"
)

// 打印1~100之间可以被7整除的值
func forDemo() {
	for i := 1; i <= 100; i++ {
		if i%7 != 0 {
			continue //结束本次循环，执行下一次循环，而break是结束整个循环。
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
}

// continue和break后也可以加标签Label。
// break带标签和不带标签的区别
func breakDemo() {
	arr1 := []string{"aaa", "bbb", "ccc"}
	arr2 := []string{"111", "222", "333"}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
		for j := 0; j < len(arr2); j++ {
			if j == 1 {
				break
			}
			fmt.Println(arr2[j])
		}
	}
}
func breakLabelDemo() {
	arr1 := []string{"aaa", "bbb", "ccc"}
	arr2 := []string{"111", "222", "333"}
LABEL:
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])

		for j := 0; j < len(arr2); j++ {
			if j == 1 {
				break LABEL
			}
			fmt.Println(arr2[j])
		}
	}

}
func continueDemo() {
	arr1 := []string{"aaa", "bbb", "ccc"}
	arr2 := []string{"111", "222", "333"}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
		for j := 0; j < len(arr2); j++ {
			if j == 1 {
				continue
			}
			fmt.Println(arr2[j])
		}
	}
}
func continueLabelDemo() {
	arr1 := []string{"aaa", "bbb", "ccc"}
	arr2 := []string{"111", "222", "333"}
LABEL:
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])

		for j := 0; j < len(arr2); j++ {
			if j == 1 {
				continue LABEL
			}
			fmt.Println(arr2[j])
		}
	}
}
func main() {
	//forDemo()
	//breakDemo()
	//breakLabelDemo()
	//continueDemo()
	continueLabelDemo()
}
