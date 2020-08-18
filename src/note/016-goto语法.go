package main

import "fmt"

/*
goto 语法可以无条件的转移到指定的行
语法格式:
goto label
...
label: xxxx;

label 泛称, 可以自定义, 只需要前后一致即可

goto 可以用于错误处理, 比如 假如有多个错误地方的处理是一致的. 那么就可以使用 goto 都指向一个地方
*/

func main() {
	for i := 1; i < 5; i++ {
		if i == 3 {
			goto test
		}
		fmt.Printf("%d \n", i)
	}
	fmt.Println("-----------这里被跳过了-----------")
test:
	fmt.Println("这里执行了")
	/**
	执行结果:
	1
	2
	这里执行了
	*/

}
