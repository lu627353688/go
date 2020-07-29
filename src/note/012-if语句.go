package main

/**
普通 if 语句

if 布尔表达式 {

} else {

}

if 布尔表达式 {

} else if 布尔表达式 {

} else {

}

增加执行体 if 语句
if 执行语句; 布尔表达式 {

}

*/

import (
	"fmt"
)

func main() {
	// 普通 if 语句
	a := 1
	if a > 0 {
		fmt.Println("a > 0")
	} else {
		fmt.Println("a <= 0")
	}

	// 添加执行体 if 语句
	if num := 10; num%5 == 0 {
		fmt.Println("0")
	} else {
		fmt.Println("1")
	}
}
