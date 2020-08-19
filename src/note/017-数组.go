package main

import "fmt"

/*
数组的定义:
	具有相同唯一类型的一组已编号且长度固定的数据项序列;
	注意:
		类型相同(类型可以是任意类型)
		长度固定
	数组的索引从 0 开始
	数组一旦定义好后, 长度不能更改

数组的声明:
var name [size] type
var name = [size]type{}
var name = [...]type{x, x, x}  不写 size 的时候, 数组的长度根据 {} 数据量长度决定
var name = []type{x, x, x}  和上面的 [...] 一样效果
var name = [size]type{index:value, ...} index可以是不连续的,下同
var name = []type{index:value, ...}
var name = [...]type{index:value, ...}
注意:
	如果在声明的时候, 没有填充值, 那么 go 语言就会生成指定数量的对应类型的零值
	{} 的长度, 或者 {} 中的 index 值, 一定不能超过 size 的大小

数组的更改
var name = [2]string{"a", "b"}
name[1] = "c"

数组的取值
var name = [2]string{"a", "b"}
var a = name[1]

数组的长度
var name = [2]string{"a", "b"}
var length = len(name)

for range 语法可以遍历数组
for index, value := range list {
	....
}

多维数组:
var name [3][3]int{}
var name = [3][3]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9}
}

数组的类型传递是值类型, 不是引用类型, 这点和python不一样要注意
即: 数组赋值给一个新变量后, 新变量就是一个新的数组, 修改新数组里面的值, 老数组并不会跟随变化

还有要注意的一点是: 在数组进行变量重新赋值的时候, 变量必须是相同类型和相同大小的. 因为数组的大小不能被改变
	所以不同长度或者不同类型的数组变量不能相互赋值. 比如:
	var a = [3]int{}
	var b = [5]int{}
	a = b  这是不被允许的, 会触发一个异常

*/

func main() {
	var aList [5]int
	aList[1] = 1
	fmt.Println(aList) // [0 1 0 0 0]

	var bList = [...]int{1, 2, 3}
	fmt.Println(bList) // [1 2 3]

	var cList = [...]int{1, 2, 3}
	fmt.Println(cList) // [1 2 3]

	var dList = []int{0: 1, 5: 6}
	fmt.Println(dList) // [1 0 0 0 0 6]

	fmt.Println(len(dList)) // 6

	for index, value := range dList {
		fmt.Printf("%d: %d,", index, value) // 0: 1,1: 0,2: 0,3: 0,4: 0,5: 6,
	}
	fmt.Println()

	var eList = [3][3]int{}
	fmt.Println(eList) // [[0 0 0] [0 0 0] [0 0 0]]

	var fList = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(fList) // [[1 2 3] [4 5 6] [7 8 9]]

	var gList = [2][2][2]int{}
	fmt.Println(gList) // [[[0 0] [0 0]] [[0 0] [0 0]]]

	var hList = [3]int{}
	var iList [3]int = hList
	iList[0] = 1
	iList[1] = 2
	iList[2] = 3
	fmt.Println(hList) // [0 0 0]
	fmt.Println(iList) // [1 2 3]
}
