package main

import (
	"fmt"

	"github.com/fatih/color"
)

var PI = 3.14159

const MAX = 100

func main() {
	// Cứ khai báo là phải sử dụng, không dùng thì tao không cho!!!
	// var email = "nguyenmanh@gmail.com"  //phải dùng
	fmt.Println("Hello Golang. Im Nguyen Manh")
	fmt.Println(PI)
	fmt.Println(MAX)
	color.Cyan("Prints text in cyan.")
}
