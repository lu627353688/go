package main

/*
显式声明常量
	const a int = 4

隐式声明常量
	const a = 4

同时声明多个常量
	const (
		a int = 1
		b int = 2
		c int = 3
	)
	在声明多个常量的时候, 如果某个常量没有设置类型和值. 那么就会继承上面的非空常量值和类型相同
	const (
		a int = 1
		b
		c
		d string = "abc"
	)
	a b c 的类型都是 int, 且 value 都是 1

注意:
	常量的类型必须是布尔型, 数字型(整数, 浮点数和复数), 字符串. 即常量只能是不可变类型(python中的解释名称)
	常量声明了可以不使用, 编译并不会报错. 这和变量不一样
*/