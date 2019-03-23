// main
package main

import (
	"fmt"
)

func main() {
	// 1 常量定义,PI是一个无类型常量
	const PI = 3.14
	fmt.Println(PI)
	// 2 也可以带类型
	const MSG_ID int = 1024
	fmt.Println(MSG_ID)
	// 3 批量声明常量
	const (
		a = 1
		b
		c
		d = 2
		e
	)
	fmt.Println(a, b, c, d, e)
	// 4 iota常量生成器
	const (
		a2 = 2 << iota
		b2
		c2
		d2
	)
	fmt.Println(a2, b2, c2, d2)

}
