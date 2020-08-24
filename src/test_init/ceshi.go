package test_init

import "fmt"

func init() {
	fmt.Println("test_init package init 函数执行了")
}

type Curl interface {
	Get()
	Post()
}

type Test struct {
}

func (t *Test) Get() {
	fmt.Println("Get")
}

func (t *Test) Post() {
	fmt.Println("Post")
}
