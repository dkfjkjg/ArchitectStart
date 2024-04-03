package main

import "fmt"

func main() {
	fmt.Println(test(30, 60))
}

func test(num1 int, num2 int) int {
	defer fmt.Println("num1:", num1)
	defer fmt.Println("num2:", num2)
	num1 += 30
	num2 += 30
	var result int = num1 + num2
	defer fmt.Println("result:", result)
	return result
}
