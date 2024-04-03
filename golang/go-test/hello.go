package main

import (
	"fmt"
	"strconv"
	"time"
)

var x, y int
var (
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

func mainHello1() {
	g, h := 123, "hello"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)

	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("面积为： %d", area)
	println()
	println(a, b, c)

	const (
		Uknown = 0
		Female = 1
		Male   = 2
	)
	println(Female)
}

func mainHello2() {
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}

func mainHello3() {
	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	fmt.Println(i, j, k, l)
}

func mainHello4() {
	var a int = 41
	var ptr *int
	ptr = &a
	println(a, *ptr)

	for true {
		fmt.Println("这是无限循环！ \n")
	}
}

func mainHello5() {
	var a, b = 100, 200
	var ret int
	ret = max(a, b)
	fmt.Println(ret)
	aa, bb := swap("Google", "Runoob")
	fmt.Println(aa, bb)
}

func swap(x, y string) (string, string) {
	return y, x
}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func mainHello6() {
	var numbers [5]int
	var numbers2 = [5]int{1, 2, 3, 4, 5}
	numbers3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers, numbers2, numbers3)

	balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)

	balance3 := [5]float32{1: 2.0, 3: 7.0}
	balance3[4] = 50
	fmt.Println(balance3)

	var n [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d \n", j, n[j])
	}
}

func mainHello7() {
	var a int = 20
	var ip *int
	ip = &a
	fmt.Printf("a 变量的地址是: %x\n", &a)
	fmt.Printf("ip变量存储的指针地址： %x \n", ip)
	fmt.Printf("*ip变量的值： %d", *ip)

	var ptr *int
	fmt.Printf("ptr的值为： %x\n", ptr)
	if ptr == nil {
		fmt.Println("ptr is nil")
	}
}

type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

func mainHello8() {
	fmt.Println(Book{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})
	fmt.Println(Book{
		title:   "Go 语言",
		author:  "www.runoob.com",
		subject: "Go 语言教程",
		book_id: 6495407,
	})
	fmt.Println(Book{
		title:  "Go 语言",
		author: "www.runoob.com",
	})
	var Book1, Book2 Book
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407
	Book2.title = "Python 语言"
	Book2.author = "www.runoob.com"
	Book2.subject = "python 语言教程"
	Book2.book_id = 6495470

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)

	printBook(Book1)

	printBookPtr(&Book2)

	var nums = make([]int, 3, 5)
	printSlice(nums)

	var numbers []int
	printSlice(numbers)
	if numbers == nil {
		fmt.Println("切片是空的")
	}
	numbers2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers2)
	fmt.Println("numbers2 == ", numbers2)
	fmt.Println("numbers2[1:4] == ", numbers2[1:4])
	fmt.Println("numbers2[2:] == ", numbers2[2:])
	fmt.Println("numbers2[:7] == ", numbers2[:7])

	var numbers3 []int
	printSlice(numbers3)

	numbers3 = append(numbers3, 0)
	printSlice(numbers3)
	numbers3 = append(numbers3, 1)
	printSlice(numbers3)
	numbers3 = append(numbers3, 2, 3, 4)
	printSlice(numbers3)

	numbers4 := make([]int, len(numbers3), cap(numbers3)*2)
	copy(numbers4, numbers3)
	printSlice(numbers4)

}

func printSlice(x []int) {
	fmt.Println("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func printBook(book Book) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func printBookPtr(book *Book) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func mainHello9() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	for key, value := range map1 {
		fmt.Printf("key is %d, value is %f\n", key, value)
	}
	for key, _ := range map1 {
		fmt.Printf("key is %d,\n", key)
	}
	for _, value := range map1 {
		fmt.Printf("value is %f\n", value)
	}

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum is: ", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func mainHello10() {
	var siteMap map[string]string
	siteMap = make(map[string]string)
	siteMap["Google"] = "谷歌"
	siteMap["Runoob"] = "菜鸟教程"
	siteMap["Baidu"] = "百度"
	siteMap["Wiki"] = "维基百科"
	for site := range siteMap {
		fmt.Println(site, "站点是 ", siteMap[site])
	}
	name, ok := siteMap["Facebook"]
	if ok {
		fmt.Println("Facebook 的 站点是", name)
	} else {
		fmt.Println("Facebook 的 站点不存在")
	}

	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	delete(countryCapitalMap, "France")

	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func mainHello11() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))

	var a int = 10
	var b float64 = float64(a)
	fmt.Println(b)
	var sum int = 17
	var count int = 5
	var mean float32
	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为 %f \n", mean)

	str := "10"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换错误！", err)
	} else {
		fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num)
	}
	str2 := strconv.Itoa(num)
	fmt.Printf("整数 %d  转换为字符串为：'%s'\n", num, str2)

	str3 := "3.14"
	num2, err := strconv.ParseFloat(str3, 64)
	if err != nil {
		fmt.Println("转换错误！", err)
	} else {
		fmt.Printf("字符串 '%s' 转换为浮点数为：%f\n", str3, num2)
	}
	str4 := strconv.FormatFloat(num2, 'f', 2, 64)
	fmt.Printf("浮点数 %f 转为字符串为：'%s'\n", num2, str4)
}

func mainHello12() {
	var i interface{} = "hello, world"
	str, ok := i.(string)
	if ok {
		fmt.Printf("%s is a string \n", str)
	} else {
		fmt.Println("conversion failed")
	}
}

type phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am a nokia, i can call you!")
}

type IPhone struct {
}

func (iphone IPhone) call() {
	fmt.Println("I am a iphone, i can call you!")
}

func mainHello13() {
	var phone phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	var s Shape
	s = Rectangle{width: 10, height: 5}
	fmt.Println("矩形面积: ", s.area())

	s = Circle{radius: 3}
	fmt.Println("圆形面积: ", s.area())
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

type Writer interface {
	write([]byte) (int, error)
}

type StringWriter struct {
	str string
}

func (sw *StringWriter) write(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}

func mainHello14() {
	var w Writer = &StringWriter{}
	sw := w.(*StringWriter)
	sw.str = "Hello,World!"
	fmt.Println(sw.str)
}

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `
   Cannot proceed, the divider is zero.
   dividee: %d
   divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(dividee int, divider int) (result int, errMsg string) {
	if divider == 0 {
		rData := DivideError{dividee: dividee, divider: divider}
		errMsg = rData.Error()
		return
	} else {
		return dividee / divider, ""
	}
}

func mainHello15() {
	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

}

func say(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(str)
	}
}

func mainHello16() {
	go say("hello")
	say("world")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, val := range s {
		sum += val
	}
	c <- sum
}

func mainHello17() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
