package main
/*
iota，特殊常量，可以认为是一个可以被编译器修改的常量

iota 可以被用作枚举值.
const (
	a = iota
	b = iota
	c = iota
)
第一个iota = 0. 当 iota 一旦被使用, 在常量组中下面每增加一行, iota 都会自动加一
所以 a = 0   b = 1   c = 2

又由于常量组在定义的时候,如果不写类型和值,就会被认为和上面的类型和值一致.
所以上面的写法可以简写为
const (
	a = iota
	b
	c
)

iota 自动加一示例
const (
	a = iota // 0
	b	// 1
	c	// 2
	d string = "test"  // test
	e	// test
	f = iota	// 5
	g	// 6
)

注意不能对数字类型进行取内存地址. 即  var a int = 1  不能进行 &a 操作
*/
import (
	. "fmt"
)

func main() {
	const (
		a = iota // 0
		b	// 1
		c	// 2
		d string = "test"  // test
		e	// test
		f = iota	// 5
		g	// 6
	)
	Println(a, b, c, d, e, f, g)
}