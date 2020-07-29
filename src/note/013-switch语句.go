package main

/**
switch 表达式 {
	case val1:
		...
	case val2:
		...
	default:
		...
}

switch 计算表达式的值与下面的每一项进行匹配
如果没有填写表达式, 那就会匹配 true, 并且每个case都会被认为是 true,即每个case都会执行
每一个 case 都默认在最后带有 break 语句
case 后面的 val 必须都是相同类型的值.
如果多个 val 会执行相同的代码块, 可以把 val 放一起,用逗号隔开,, 比如 case val1, val2, val3:
如果想要执行完匹配的case后继续执行下面的case,可以使用 fallthrough, 并且 fallthrough 只能是每个case的最后一行
	注意: fallthrough 只会击穿一次. 如果想要一直执行下次, 需要在后面的case一直写 fallthrough

*/

import (
	"fmt"
)

func main() {
	a := 1
	switch a {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	default:
		fmt.Println("undefined")
	}
	/**
	执行结果: 1
	*/
	switch a {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	default:
		fmt.Println("undefined")
	}
	/**
	执行结果 1 2
	*/
	switch a {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
		fallthrough
	case 3:
		fmt.Println(3)
		fallthrough
	case 4:
		fmt.Println(4)
		fallthrough
	default:
		fmt.Println("undefined")
	}
	/**
	执行结果 1 2 3 4 undefined
	*/
}
