package main

/*
func Scan(a ...interface{}) (n int, err error)

func Scanf(format string, a ...interface{}) (n int, err error)

func Scanln(a ...interface{}) (n int, err error)

*/

import (
	"fmt"
)

func main() {
	var x int
	var y float64
	fmt.Println("请输入一个整数，一个浮点类型：")
	fmt.Scanln(&x, &y) //读取键盘的输入，通过操作地址，赋值给x和y   阻塞式
	fmt.Printf("x的数值：%d，y的数值：%f\n", x, y)

	fmt.Scanf("%d,%f", &x, &y)
	fmt.Printf("x:%d,y:%f\n", x, y)
}

/**
请输入一个整数，一个浮点类型：
11 1.2
x的数值：11，y的数值：1.2
1,2
x:1,y:2
*/
