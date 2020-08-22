package main

import "fmt"

/**
在函数内有一种语法叫 defer (延迟)
在函数结束之前会执行 defer 后面的语句
如果一个函数内声明了多行defer语句. 那么会逆序执行, 即从下到上依次执行
报错也不会影响 defer 语句的运行.

注意:
	如果 defer 后面跟的是一个函数的运行, 并且函数传递了参数.
	那么函数传递的参数会锁定, 但是函数里面的执行体不会立即执行, 而是在最后执行
	具体看下面示例
*/

func test(x int, y int) {
	fmt.Printf("x point: %p, y point: %p \n", &x, &y)
	fmt.Printf("x: %d, y: %d \n", x, y)
}

func test_point_arg(_slice *[]int) {
	fmt.Printf("test_point_arg _slice point: %p \n", _slice)
	fmt.Printf("test_point_arg _slice[0]: %d \n", (*_slice)[0])
}

func main() {
	a := 1
	b := 2
	fmt.Printf("a: %d, b: %d \n", a, b)
	fmt.Printf("a point: %p, b point: %p \n", &a, &b)
	defer test(a, b)
	a = 100
	b = 200
	fmt.Printf("a: %d, b: %d \n", a, b)
	fmt.Printf("a point: %p, b point: %p \n", &a, &b)
	/**
	输出结果:
		a: 1, b: 2
		a: 100, b: 200
		x: 1, y: 2
	虽然 test 函数是最后执行的, 但是执行到 defer test(a, b) 这行的时候 a, b 的值会以当前时刻的值锁定,
	后面无论怎么修改 a, b 最后执行的时候都不会影响结果.
	当然也有例外. 传递指针的时候(引用传递), 如果修改了指针指向内存中的值, 结果依然会受到影响.
	但是这个影响并不是就打破了上面的值锁定问题, defer 语句声明的时候参数依然是锁定的, 只不过锁定的值是指针的时候比较特殊
	这里面牵涉到 值传递 和 引用传递.
	第一个例子是值传递, defer 声明的时候是复制了一份数据, 可以看到函数 test 内执行的时候, 地址和外面已经不一样了
	看下面的例子
	*/
	_slice := []int{0}
	fmt.Printf("_slice[0]: %d \n", _slice[0])
	fmt.Printf("_slice point: %p \n", &_slice)
	defer test_point_arg(&_slice)
	_slice[0] = 1
	fmt.Printf("_slice[0]: %d \n", _slice[0])
	fmt.Printf("_slice point: %p \n", &_slice)
	/**
	输出结果:
		a: 1, b: 2
		a point: 0xc0000100a8, b point: 0xc0000100c0
		a: 100, b: 200
		a point: 0xc0000100a8, b point: 0xc0000100c0
		_slice[0]: 0
		_slice point: 0xc000004460
		_slice[0]: 1
		_slice point: 0xc000004460
		test_point_arg _slice point: 0xc000004460
		test_point_arg _slice[0]: 1
		x point: 0xc000010128, y point: 0xc000010130
		x: 1, y: 2
	上面两次输出结果对比可以验证:
		1. 有多个 defer 语句的时候, 最后声明的优先执行. 先声明的后执行
		2. defer 后面跟随的函数调用, 所依赖的参数并不是最后执行函数执行体的时候才去索引值的. 而是声明的时候就锁定了
			注意值传递和引用传递
	*/
}
