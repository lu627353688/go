package main

import "fmt"

/**
map集合将 key 和 value 对应起来, 可以快速的检索数据

注意:
	map 和 python dict 一样是无序的, 且是引用传递, 注意数组是值传递, 切片是引用传递
	map 和 切片一样, 长度是不固定的. 数组的长度是固定的
	len 函数 一样可以对 map 使用来获取长度( key 的数量).
	所有可比较的类型都可以当做 key, 比如 布尔型 整数型 浮点型 字符串 复合类型

map 的创建(必须使用 make, 下面两步缺一不可)
var name map[key_type]value_type  // 定义结构
var name = make(map[key_type]value_type)  // make 初始化
如果是在函数内可以简写:
name := map[key_type]value_type{"key": value, ...} // 结构和初始化一步到位

注意:
	map 一定要初始化, 否则不能存放键值对. 比如以下情况就不能存放键值对
	var test map[string]string  这种情况只是声明了 map 的结构, 后面没有 {"key": value} 相当于是没有初始化的
	test["haha"] = "haha" 然后会报错, nil map 不能赋值.
	想要正确使用, 要下面这么写:
	方式一:
	var test map[string]string  定义结构(创建 nil map)
	test = make(map[string]string) 初始化 map
	test["haha"] = "haha" 赋值
	方式二: 一步到位
	test := map[string]string{"haha": "haha"}

map 一样可以使用 range 进行遍历
语法:
for key := range map {}

map 赋值
name[key] = value

map 取值
value, ok = name[key]
注意:
如果 key 在 name 中不存在, 是不会报错的. 会返回 map 设置的 value 对应类型的空值(比如字符串类型是 "").
ok 代表 key 是否在 map 中存在, 布尔类型

map 删除值 (不返回任何值), 删除不存在的 key 也不会报错
delete(map_name, key)


*/

func main() {
	test := map[string]string{"haha": "haha"}
	fmt.Println(test)
	for key := range test {
		fmt.Println(key, test[key])
	}
	delete(test, "no_exist_key")
}
