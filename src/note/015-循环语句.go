package main

import "fmt"

/*
在 go 语言中, for 是唯一的循环, 没有 while 循环
语法结构:
	for init; condition; exec {}
		init: 初始化语句, 只会执行一次
		condition: 条件语句, 在初始化语句执行完之后, 就会判断该语句, 如果结果为 true, 就会执行一次结构体语句 {}
		exec: 每次成功循环之后({}执行完后)都会执行的语句. 执行完该语句之后,会继续检查条件, 如果条件为 true, 会再次执行 {} 如果为 false, 退出循环

注意:
	init 和 condition 和 exec 都是可选的.即
	for condition {} 是允许的
	for {} 也是允许的 (不建议这么写)

for range 语法 可以对 slice, map, 数组, 字符串等进行迭代循环
比如:
for key, value := range map1 {
	map2[key] = value
}

跳出循环语法: break
跳过当前循环执行下一次循环: continue
*/

func main() {
	for i := 1; i < 10; i++ {
		fmt.Printf("%d", i)
	}
	var a, b = 6, 10
	for a < b {
		fmt.Printf("%d", a)
		a++
	}
	var intS = [5]int{1, 2, 3, 4, 5}
	for index, value := range intS {
		fmt.Printf("第%d个数字是%d", index, value)
	}
}
