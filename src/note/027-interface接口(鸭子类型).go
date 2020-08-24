package main

/**
接口定义了一系列方法
如果某个类型想要 实现接口, 就要实现接口所定义的所有函数

真正核心作用: 接口变量可以接收实现接口的所有类型对象, 那么写函数或者方法的时候, 就可以让函数或者方法接收任意类型的变量

接口的定义语法:
type interface_name interface {
	method_name1(arg type) (res type)
	method_name2(arg type) (res type)
	...
}

接口的变量声明
var i interface_name

注意: 空接口可以接受任意类型, 因为空接口没有定义任何方法
type empty_interface interface {

}
*/

import (
	"test_init"
)

type request struct {
	name string
	test_init.Test
}

func main() {
	var r *request = new(request)
	r.name = "a"
	var i test_init.Curl
	i = r
	i.Get() // Get
	r.Get() // Get
}
