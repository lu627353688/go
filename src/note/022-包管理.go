package main

/**
go 语言代码包 package 是 src 目录. src 目录下的每一个目录都是一个包
同一个包下面的每一个文件的第一行的 package 名字都必须一致
包名应和包下面的每个文件的第一行的 package 名称一致(也可以不一致, 但是最好都一致)
名字为 main 的包是程序的入口

同一个包下面的所有文件属于同一个工程, 不需要引入就可以使用

需要被导出的函数和常量, 需要以大写字母开头

import 的几种操作:
1. 普通操作
import "fmt"  // 单个包导入
import (  // 多个包导入
	"fmt",
	"xxxx",
)

2. . 操作
import (
	. "fmt"
)
这个操作是使用的时候可以省略包名. 比如 fmt.Println("hello world") 如果按照上面的 . 方式导入就可以写为 Println("hello world")

3. 别名操作
import (
	_fmt "fmt"
)
fmt.Println("hello world") 就要写为 _fmt.Println("hello world")

4. _ 操作 (不导入包, 只执行包的初始化操作( init 函数))
import (
	_ "github.com/ziutek/mymysql/godrv"
)
对于 go 语言来说 main 函数和 init 函数都是 保留函数. 在 import 包的时候会优先执行 init 函数
所以 init 函数会优先 main 函数执行


main 函数 和 init 函数的注意点:
	两个函数在定义时不能有任何的参数和返回值。 该函数只能由 go 程序自动调用，不可以被引用。
	init 可以应用于任意包中，且可以重复定义多个。 main 函数只能用于 main 包中，且只能定义一个。
	两个函数的执行顺序：
		在 main 包中的 go 文件默认总是会被执行。
		对同一个 go 文件的 init( ) 调用顺序是从上到下的。
		对同一个 package 中的不同文件，将文件名按字符串进行“从小到大”排序，之后顺序调用各文件中的init()函数。
		对于不同的 package，如果不相互依赖的话，按照 main 包中 import 的顺序调用其包中的 init() 函数。
		如果 package 存在依赖，调用顺序为最后被依赖的最先被初始化，例如：导入顺序 main –> A –> B –> C，则初始化顺序为 C –> B –> A –> main，一次执行对应的 init 方法。main 包总是被最后一个初始化，因为它总是依赖别的包
	为避免出现循环 import, 一个包被多个包引用的时候, 只会初始化一次

下载包命令:
	go get package
	会自动安装在 gopath 设置的目录下面的 src 目录中
*/

import (
	"fmt"
	_ "test_init"
)

func main() {
	fmt.Println("main 函数启动了")
	/**
	输出结果:
		test_init package init 函数执行了
		main 函数启动了
	*/
}
