package main

import (
	"fmt"
)

func main() {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型 :%T", i)
	case int:
		fmt.Println("x 是 int 类型")
	case float64:
		fmt.Println("x 是 float64 类型")
	case bool, string:
		fmt.Println("x 是 bool 类型, 或者 string 类型")
	}
}
