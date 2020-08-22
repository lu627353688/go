package main

import "fmt"

/**
函数的语法:
func funcName (arg1 type1, arg2 type2) (res1 type1, res2 type2) {
	return a, b
}
注意:
	res1 和 res2 这两个变量名字可以不声明, 但是类型是一定要声明的
	当只有一个返回值的时候可以不声明变量也不声明类型, 如果都不声明括号也就不用写了

变长参数:
	func funcName (args ...string) {}
	// args ...string 指定了这个函数接受不定数量的参数. 这个 args 会处理成一个切片

函数传参如果想要引用传递必须要使用指针. 比如传递 int_var_name, 要换成 &int_var_name , 接收参数的地方改成 *int
注意: slice 和 map 最好也传递指针. 因为 slice 可能会改变长度(扩容)

局部变量:
	在函数内声明的变量叫做局部变量

全局变量:
	在全局定义的变量叫全局变量

*/

func test(x int, y int) (m int, n int) {
	m = 2 * (x + y)
	n = x * y
	return m, n
}

func test_args(args ...int) {
	for index, value := range args {
		fmt.Println(index, ": ", value)
	}
}

func main() {
	a := 1
	b := 2
	fmt.Println(test(a, b)) // 6 2
	test_args(1, 2, 3, 4)
	/*
		0 :  1
		1 :  2
		2 :  3
		3 :  4
	*/
}
