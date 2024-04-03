package main

import "fmt"

func main1() {
	fmt.Println("Hello,World!")
	fmt.Println("Google" + "Runoob")
	var stockcode = 123
	var enddate = "2020-12-13"
	var url = "code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
	var a string = "Runoot"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	var bbool bool
	fmt.Println(bbool)

	var i int
	var f float64
	var bbool2 bool
	var ss string
	fmt.Printf("%v %v %v %q\n", i, f, bbool2, ss)

	ff := "runoob"
	fmt.Println(ff)
}
