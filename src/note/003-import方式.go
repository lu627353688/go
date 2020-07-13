package main

/*
1.点操作 我们有时候会看到如下的方式导入包
import(
	. "fmt"
)
这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调
用的fmt.Println("hello world")可以省略的写成Println("hello world")


2.别名操作 别名操作顾名思义我们可以把包命名成另一个我们用起来容易记忆的名字
import(
	f "fmt"
)
别名操作的话调用包函数时前缀变成了我们的前缀，即f.Println("hello world")


3._操作
import (
  "database/sql"
  _ "github.com/ziutek/mymysql/godrv"
)
_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数
*/
