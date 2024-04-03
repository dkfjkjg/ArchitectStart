package testutils

import "fmt"

var Age int
var Sex string
var Name string

func init() {
	fmt.Println("testutils中的init函数被执行了")
	Age = 10
	Sex = "女"

	Name = "lili"
}
