package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type Accstatus string

type Student struct {
	Name   string
	Email  string
	Age    int
	Gender *string

	Status Accstatus
}

func Sum(number ...int) int {
	return 10
}

type Email func(string, string, string)

func main() {
	var student Student = Student{
		Name:  "Manh",
		Email: "nguyenmanh@gmail.com",
		Age:   22,
	}
	spew.Dump(student)
	fmt.Println(student)

	// anonymous struct??
	// ứng dụng cho unit test, khi muốn khai báo tường mình giá trị
	var demoStr = struct {
		Name  string
		Email string
	}{
		Name:  "Manh",
		Email: "nguyena@gmail.com",
	}

	spew.Dump(demoStr)
}
