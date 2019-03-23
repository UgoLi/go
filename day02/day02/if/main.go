// main
package main

import (
	"fmt"
)

func ifDemo() {
	fmt.Println("请输入您的成绩(0~100)：")
	var score int
	fmt.Scan(&score)
	if score < 0 || score > 100 {
		fmt.Println("您输入的成绩有误！")
		return
	}
	if score >= 60 {
		fmt.Println("恭喜您，通过考试！")
	} else {
		fmt.Println("很遗憾，没有通过考试！")
	}
	if score >= 90 {
		fmt.Println("您的成绩为A!")
	} else if score < 90 && score >= 80 {
		fmt.Println("您的成绩为B!")
	} else if score < 80 && score >= 60 {
		fmt.Println("您的成绩为C!")
	} else {
		fmt.Println("您的成绩为D!")
	}

}

// 根据输入的年份和月份，输出月份的天数，需要判断闰年(15分钟)
func switchDemo() {
	fmt.Println("请输入年份月份：")
	var year, month int
	fmt.Scan(&year, &month)
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31天")
	case 4, 6, 9, 11:
		fmt.Println("30天")
	case 2:
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			fmt.Println("29天")
		} else {
			fmt.Println("28天")
		}
	default:
		fmt.Println("输入有误！")
	}
}

// switch后不加条件，每个case后加条件。
func switchDemo2() {
	fmt.Println("请输入您的成绩(0~100):")
	var score int
	fmt.Scan(&score)
	switch {
	case score <= 100 && score >= 90:
		fmt.Println("成绩等级为A")
	case score < 90 && score >= 80:
		fmt.Println("成绩等级为B")
	case score < 80 && score >= 60:
		fmt.Println("成绩等级为C")
	case score >= 0 && score < 60:
		fmt.Println("成绩等级为D")
	default:
		fmt.Println("成绩有误")
	}
}

func main() {
	//ifDemo()
	//switchDemo()
	switchDemo2()
}
