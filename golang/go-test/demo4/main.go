package main

import "fmt"

func main() {
	test()
	fmt.Println("方法后执行")
}

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("报错信息,", err)
		}
	}()
	var num1 int = 10
	var num2 int = 0
	var result int = num1 / num2
	fmt.Println(result)
}
