// main
package main

import (
	"fmt"
)

func mapDemo1() {
	// 1 声明映射
	var m1 map[string]int
	fmt.Println(m1)
	fmt.Println(m1 == nil)
	// 2 创建，cap
	m1 = make(map[string]int, 10)
	// 3 赋值
	m1["张三"] = 90
	m1["李四"] = 95
	m1["王五"] = 80
	m1["刘六"] = 55
	fmt.Println(m1)
	// 4 删除
	delete(m1, "李四")
	delete(m1, "杨琪")
	fmt.Println(m1)
	// 5 查找元素
	if v, ok := m1["刘六"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("键为刘六的元素不存在")
	}
	// 6 遍历映射
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 7 映射不能做取地址操作
	//fmt.Println(&m1["张三"])//error

}
func main() {
	mapDemo1()
}
