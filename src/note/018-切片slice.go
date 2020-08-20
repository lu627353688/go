package main

import "fmt"

/**
数组的大小是固定的, 不能进行动态扩容, 这在很多情况下是非常不方便的,
所以go语言提供了另外一种数据类型切片(也叫动态数组)
切片的底层一定会有一个数组. 切片是不存储数据的, 数据都在数组中存储.
所以切片的数据如果改变了, 数组中的数据也会改变

切片的组成有三部分:
	指针 指向开始位置
	长度 切片的长度  len
	最大长度 开始位置到最后位置的长度  cap 也叫容量

切片的创建:
var name = []type

也可以用make创建
var name []type = make([]type, len, cap)
如果是在函数内部可以简写为:
name := make([]type, len)

从切片创建一个新切片
new := old[s_index:e_index]  左闭右开区间
如果 s_index 省略代表从头开始
如果 e_index 省略代表截取到最后
切片的长度为 e_index - s_index

添加元素 append 方法
append 之后如果切片的长度将要大于切片的容量后, 切片会进行扩容. 切片的容量总是偶数 2 * n
每扩容一次 (n + 1) 所以应该是这样的 0 2 4 6 8 10, 而不是 0 2 4 8 16 32 64. 这点需要注意

如果你想要复制一个切片, 改变新切片不影响老的切片应该用 copy 方法
copy(new, old)  会从老的切片中复制数据到新的切片中. 新的切片可以按照上面的 make 方法创建
*/

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func main() {
	var numbers []int
	printSlice(numbers) // len=0 cap=0 slice=[]

	numbers = append(numbers, 0)
	printSlice(numbers) // len=1 cap=1 slice=[0]

	numbers = append(numbers, 1)
	printSlice(numbers) // len=2 cap=2 slice=[0 1]

	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers) // len=5 cap=6 slice=[0 1 2 3 4]

	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	copy(numbers1, numbers)
	printSlice(numbers1) // len=5 cap=12 slice=[0 1 2 3 4]

	numbers1[0] = 1      // 新的切片不会影响老的切片
	printSlice(numbers1) // len=5 cap=12 slice=[1 1 2 3 4]
	printSlice(numbers)  // len=5 cap=6 slice=[0 1 2 3 4]
}
