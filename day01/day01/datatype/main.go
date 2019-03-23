// main
package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

// 1 整数
func IntegerDemo() {
	var a int = 100
	// 1 我想知道a的长度，占几个字节
	fmt.Println(unsafe.Sizeof(a))
	// 2 ++、--是语句，不是表达值;++只能在变量后
	x := 10
	//y:=x++ //error
	x++
	//++x //++只能在变量后
	y := x
	fmt.Println(x, y)
	// && ||的短路运算
	// f(x)||g(y),如果f(x)的结果为真，不再计算g(y)。
	// f(x)&&g(y),如果f(x)的结果为假，也不再计算g(y)。
	// 位运算
	x1 := 2
	y1 := 15
	fmt.Println(x1&y1, x1|y1, x1^y1, x1&^y1)
}

// 2 浮点数
func floatDemo() {
	var f1 float32 = 12.5
	fmt.Println(f1)
	f2 := 22.6
	fmt.Println(f2, reflect.TypeOf(f2))
	// 1 float32与float64之间需要强制转换
	f2 = float64(f1)
	fmt.Println(f2)
	// 2 浮点数与零比较:浮点数的绝对值小于一个极小的数
	// 3 科学计数法
	f3 := 1.99e+5
	f4 := 1.99E-3
	fmt.Println(f3, f4)

}

func intDemo() {
	//只要类型不相同，都需要强制转换
	var a int8 = 16
	var b int32 = 0
	fmt.Println(a, b)
	b = int32(a)
	fmt.Println(b)

}

// 4复数
func complexDemo() {
	c1 := 10 + 5i
	fmt.Println(c1, reflect.TypeOf(c1))
	var c2 complex64 = 1.2 + 7.8i
	fmt.Println(c2, reflect.TypeOf(c2))
	fmt.Println(c1 + complex128(c2))
	c3 := 1.6 - 2.5i
	fmt.Println(c1+c3, c1-c3, c1*c3)
	// 分别打印c3的实部和虚部
	fmt.Println(real(c3), imag(c3))

}

// 5 布尔类型
func boolDemo() {
	flag := false
	fmt.Println(flag)
	var flag2 bool
	fmt.Println(flag2)
	flag2 = true
	fmt.Println(flag2)
}

// 6 字符串类型
func stringDemo() {
	s1 := "Hello"
	s2 := " 世界"
	// 1 字符串连接
	s3 := s1 + s2
	fmt.Println(s3)
	// 2.1 字符串长度（以字节为单位）
	fmt.Println(len(s3))
	// 2.2 包含的字符数量(以字符为单位)
	fmt.Println(utf8.RuneCountInString(s3))
	// 3 截取子串(包含s3[1],不包含s3[5],左闭右开)
	fmt.Println(s3[1:5])
	fmt.Println(s3[:])
	fmt.Println(s3[:3])
	fmt.Println(s3[3:])
	fmt.Println(s3[0]) //打印的是字母的asc码值
	//s3[0] = "h"        //不允许
	// 4 遍历
	// 4.1 字节遍历
	for i := 0; i < len(s3); i++ {
		fmt.Print(s3[i], " ")
	}
	fmt.Println()

	// 4.2 字符遍历
	for _, ch := range s3 {
		fmt.Print(string(ch), " ")
	}
	fmt.Println()
	// 5 原生字符串字面值
	str1 := `  静夜思 
窗前明月光
疑是地上霜 
举头望明月 
低头思故乡`
	fmt.Println(str1)
	// 6 unicode与utf-8，utf-16之间关系？
	// 6.1 unicode就是一种字符编码。
	// 6.2 utf-8和utf-16是两种不同的对unicode编码的存储方式
	// 7 字符串比较
	s4 := "Hello 世界"
	fmt.Println(s3 == s4) //true
}

// 7 指针类型
func ptrDemo() {
	var num int = 100
	var p *int = &num
	// 二级指针
	var pp **int = &p
	fmt.Println(num, *p, **pp)
	fmt.Println(&num, p, *pp)
}

type Age int   //年龄类型
type Score int //分数类型

// 8 type定义类型
func typeDemo() {
	var myAge Age = 35
	fmt.Println(myAge)
	var myScore Score = 116
	fmt.Println(myScore)
	myAge = Age(myScore)
	fmt.Println(myAge)
}

// 包级别的
var Num int = 1000

// 9 变量的作用域
func regionDemo() {
	var num2 int = 100 //num的作用域是整个函数
	if num2 == 100 {
		num3 := 20 //语法块级别的变量，num3的作用域if{}块
		fmt.Println(num3)
	}
	fmt.Println(Num)
	fmt.Println(num2)
	//fmt.Println(num3)//超出作用域
}

// 变量重名问题，就近原则
func regionDemo2() {
	fmt.Println(Num) //包级别的
	var Num int = 20

	if true {
		var Num int = 10 //语法块级别的
		fmt.Println(Num)
	}
	fmt.Println(Num) //函数级别的
}
func main() {

	//IntegerDemo()
	//floatDemo()
	//intDemo()
	//complexDemo()
	//boolDemo()
	//stringDemo()
	//ptrDemo()
	//typeDemo()
	//regionDemo()
	regionDemo2()
}
