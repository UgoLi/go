// main
package main

import (
	"fmt"
)

// 输出,控制台打印
func Output() {
	// 1 输出，并在末尾自动增加换行
	fmt.Println("Hello World")
	// 2 输出，在末尾不会自动增加换行
	fmt.Print("Hello World\n")
	// 3 格式化输出，也会自动增加换行
	name := "张三"
	age := 18
	fmt.Printf("姓名:%s,年龄:%d\n", name, age)
	// 4  格式化并返回字符串
	str := fmt.Sprintf("姓名:%s,年龄:%d\n", name, age)
	fmt.Println(str)
}

// 输入，接收控制台中用户输入的值
func Input() {
	fmt.Println("请输入两个整数:")
	var a, b int
	// 1 变量之间的分隔符可以是空格、tab键，回车键
	//fmt.Scan(&a, &b)
	// 2 变量之间的分隔符可以是空格、tab键
	//fmt.Scanln(&a, &b)
	// 3 变量之间的分隔符在输入时，要和第一参数的分隔符相同
	fmt.Scanf("%d;%d", &a, &b)
	fmt.Println(a, b)

}
func main() {
	//Output()
	Input()
}
