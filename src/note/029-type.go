package main

/**
type 不仅能定义 结构体 struct 接口 interface 还可以定义其他类型

语法:
type 定义的类型名字 已有的type名字

type myInt int
type myStr string
type myFunc func(arg type)(res type)

注意: 上面这种定义新类型后. 新类型就不是原类型. 即 myInt 变量不能接收 int 型变量.
但是变量别名可以(需要 go 是 1.9 之后的版本)

变量别名:
type myInt = int
这种是起一个别名, 但实质还是他自己. myInt 类型变量就可以接收 int 类型变量

不能对 go 语言 本身已有的类型 别名后 添加新的方法
自己定义的本地类型可以. 但是如果是go本身定义的就不行
*/

import "fmt"

type myint int
type mystr string
type myFunc func(arg1 string) (res1 string)
type eInt = int

func getMyFunc() (f myFunc) {
	f = func(arg1 string) (res1 string) {
		res1 = arg1
		return res1
	}
	return f
}

func main() {

	var i1 myint
	var i2 = 100
	i1 = 100
	fmt.Println(i1)
	//i1 = i2 // 执行该语句会触发异常: cannot use i2 (type int) as type myint in assignment
	fmt.Println(i1, i2)

	var name mystr
	name = "王二狗"
	var s1 string
	s1 = "李小花"
	fmt.Println(name)
	fmt.Println(s1)
	//name = s1 // 执行该语句会触发异常:cannot use s1 (type string) as type mystr in assignment

	fmt.Println(getMyFunc()("hello world")) // hello world

	var x int = 10
	var y eInt = 20
	fmt.Println(x, y) // 10 20
	y = x
	fmt.Println(x, y) // 10 10
}
