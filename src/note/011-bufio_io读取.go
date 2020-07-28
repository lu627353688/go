package main

/**
https://golang.google.cn/pkg/bufio/
*/
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("请输入一个字符串：")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	fmt.Println("读到的数据：", s1)
}

/**
请输入一个字符串：
hello world
读到的数据：hello world
*/
