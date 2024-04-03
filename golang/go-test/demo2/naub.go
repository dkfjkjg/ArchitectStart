package main

import "fmt"

var func01 = func(num1 int, num2 int) int {
	return num1 * num2
}

func main() {
	result := func(num1 int, num2 int) int {
		return num1 + num1
	}(10, 20)
	fmt.Println(result)

	sub := func(num1 int, num2 int) int {
		return num1 - num2
	}

	result01 := sub(30, 70)
	fmt.Println(result01)

	result02 := sub(30, 70)
	fmt.Println(result02)

	result03 := func01(3, 4)
	fmt.Println(result03)
}
