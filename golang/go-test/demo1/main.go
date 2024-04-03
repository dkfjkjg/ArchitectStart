package main

import (
	"fmt"
	"go-test/demo1/dbutils"
	"go-test/demo1/testutils"
)

var num int = test()

func test() int {
	fmt.Println("test函数被执行！")
	return 10
}

func init() {
	fmt.Println("init函数被执行！")
}

func main() {
	fmt.Println("main函数被执行！")
	fmt.Println("Age=", dbutils.Age, "Sex=", testutils.Sex, "Name=", testutils.Name)
	fmt.Println("Age=", testutils.Age, "Sex=", testutils.Sex, "Name=", testutils.Name)
}
