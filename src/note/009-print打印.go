package main

/*
打印:
func Print(a ...interface{}) (n int, err error)

格式化打印:
func Printf(format string, a ...interface{}) (n int, err error)

打印后换行:
func Println(a ...interface{}) (n int, err error)

格式化打印中的常用占位符：
	%v,原样输出
	%T，打印类型
	%t,bool类型
	%s，字符串
	%f，浮点
	%d，10进制的整数
	%b，2进制的整数
	%o，8进制
	%x，%X，16进制
		%x：0-9，a-f
		%X：0-9，A-F
	%c，打印字符
	%p，打印地址

*/
import (
	"fmt"
)

func main() {
	a := 10
	b := 3.14
	c := true          // bool
	d := "Hello World" //string
	e := `Ruby`        //string
	f := 'A'
	fmt.Printf("%T,%b\n", a, a)
	fmt.Printf("%T,%f\n", b, b)
	fmt.Printf("%T,%t\n", c, c)
	fmt.Printf("%T,%s\n", d, d)
	fmt.Printf("%T,%s\n", e, e)
	fmt.Printf("%T,%d,%c\n", f, f, f)
	fmt.Println("-----------------------")
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
}
