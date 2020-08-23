package main

import "fmt"

/**
在 go 语言中, 数组只能存储相同类型的数据. 如果我们想要存储不同类型数据集怎么办? 解决办法就是结构体
结构体定义:
结构体是由一系列具有相同类型或不同类型的数据构成的数据集合

结构体语法:
type struct_name struct {
	property1 type1
	property2 type2
	...
}

结构体定义好之后就可以使用了.

赋值方式:
假设有结构体:
type students struct {
	name string
	age int
}
方式一:按照顺序赋值
var mrHong students = students{"mrHong", 18}

方式二:按照 key: value 方式赋值, 可以乱序
var mrHuang students = students{age: 18, name:"mrHuang"}

方式三:new // 所有属性都会被设置为对应类型的初始值.
var mrLu students = new(students)
mrLu.name = "mrLu"
mrLu.age = 18

结构体属性的读取使用 . 方式

结构体的继承:
	结构体里面可以嵌套结构体.
	如果一个结构体内部嵌套了其他结构体, 那么就自动拥有了嵌套的结构体的所有属性

type person struct {
	name string
	age int
}

type student struct {
	person
	class string
}
这样继承了之后 student 也有 name 和 age 两个属性了.
继承之后, 字段的等级会被提升. 即:
lu := new(student)
既可以这样赋值.  lu.name = "lu"
也可以这样赋值.  lu.student.name = "lu"
注意: 如果嵌套的结构体和本身属性重复了, 那么会优先使用本身的属性.

结构体和属性的导出
如果结构体想要导出, 允许其他包导入后使用, 需要大写首字母
结构体的属性也是同理, 只有大写字母开头的属性才可以被其他包使用
*/

type person struct {
	name string
	age  int
}
type student struct {
	person
	class string
}

func main() {
	lu := new(student)
	lu.name = "lu"
	lu.person.age = 18
	lu.class = "xxx系xxx班"
	fmt.Println(lu) // &{{lu 18} xxx系xxx班}
}
