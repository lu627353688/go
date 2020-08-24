package main

import "fmt"

/**
go 语言中同时有函数和方法.
一个方法就是一个包含了接受者的函数, 接受者可以是命名类型或者结构体类型的一个值或者是一个指针.
所有给定类型的方法属于该类型的方法集

函数语法:
func func_name (args type) (res type) {

}

方法语法:
func (receiver type) func_name (args type) (res type) {

}

不同的 receiver 可以拥有相同的 func_name , 互不影响

结构体嵌套结构体的时候,方法也是可以被继承的, 和 属性一样, 重复了都是采用就近原则.
上面这一行为 也叫重写方法

可以看出 这种方式 其实就是模拟实现 面向对象 类

注意, 在 go 语言中, 结构体在赋值给另一个变量的时候也是值传递, 相当于重新赋值了一份

*/

type person struct {
	name string
	age  int
}

type student struct {
	person
	class string
}

func (x person) speak() {
	fmt.Printf("x.name: %s, x.age: %d \n", x.name, x.age)
}

func (x person) modify() {
	x.age = 20
}

func (x *person) modify2() {
	x.age = 30
}

func main() {
	a := person{"a", 18}
	b := a
	fmt.Printf("a point: %p \n", &a) // a point: 0xc000004440
	fmt.Printf("b point: %p \n", &b) // b point: 0xc000004460

	a.speak() // x.name: a, x.age: 18
	a.modify()
	a.speak() // x.name: a, x.age: 18
	a.modify2()
	a.speak() // x.name: a, x.age: 30
	// 由上面也可以看出, 不设置接受对象是指针的话, 在函数内的修改不会影响原则,因为是值传递
	// 还有一个有趣的是, 虽然你设置了接收者是指针, 但是你不用写 &a.modify2()  这是因为 go 语言做了处理

	c := student{person{"c", 18}, "xxx"}
	c.speak() // x.name: c, x.age: 18
}
