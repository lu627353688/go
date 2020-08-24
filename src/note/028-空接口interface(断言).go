package main

import "fmt"

/**
断言 语法格式: 断言成功,返回值就是真正的值, flag 为true. 如果断言失败, 那么flag为false, 返回值为断言的类型的零值
安全类型断言: 断言失败不会引发恐慌, 布尔值为 false.
返回值, 布尔值 := 变量名.(断言类型)

非安全类型断言: 如果断言失败会触发异常
返回值 := 变量名.(断言类型)

*/

type empty_interface interface {
}

type person struct {
	name string
	age  int
}

func main() {
	p := person{"test", 18}
	var i empty_interface
	i = p

	a, flag1 := i.(person)
	b, flag2 := i.(int)
	fmt.Println("i 是person? ", flag1, ", a 值是:", a)
	fmt.Println("i 是int? ", flag2, ", b 值是:", b)
	/**
	输出结果:
	i 是person?  true , a 值是: {test 18}
	i 是int?  false , b 值是: 0

	由上面可以看出:
	第一次因为断言成功了, 所以返回值是原本的值 person{"test", 18}
	第二次断言失败了, 所以 flag 为 false, 返回的值是 断言类型是 int 的零值
	*/
}
