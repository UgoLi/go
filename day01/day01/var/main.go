// main
package main

import (
	"fmt"
	"reflect" //反射包
)

// 函数外的变量称为包级别的变量，包级别的变量不能使用简短声明
//a:=100
var a int = 100

func main() {
	// 1 变量的一般声明格式
	var num int = 100
	fmt.Println("num=", num)
	// 2 简短声明,自动推断类型
	age := 18
	fmt.Println("age=", age)
	// age这个变量的类型
	fmt.Println("age的类型是：", reflect.TypeOf(age))
	fmt.Println("a=", a)
	// 3 变量没有赋值，默认会设置为该类型的 "零" 值
	var b int
	fmt.Println("b=", b)
	// 4 多个变量声明
	var a1, a2, a3 int
	fmt.Println(a1, a2, a3)
	var b1, b2, b3 = 10, 12.5, "Hello"
	fmt.Println(b1, b2, b3)
	name, age := "张三", 18
	fmt.Println("name=", name, "age=", age)
	var (
		bookid   int
		bookname string
	)
	fmt.Println(bookid, bookname)
	// 5 多重赋值
	x, y := 10, 20
	fmt.Println(x, y)
	// 交互x与y的值
	x, y = y, x
	fmt.Println(x, y)
	// _标识符
	_, x1 := MultiReturnValue()
	fmt.Println(x1)
	s1, _ := MultiReturnValue()
	fmt.Println(s1)
}
func MultiReturnValue() (string, int) {
	return "Hello", 100
}
