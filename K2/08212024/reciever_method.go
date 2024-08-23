package main

import "fmt"

// Khai báo một struct
type Rectangle struct {
	width, height int
}

// Phương thức với receiver theo giá trị
func (r Rectangle) area() int {
	return r.width * r.height
}

func main() {
	rect := Rectangle{width: 10, height: 5}
	fmt.Println("Area:", rect.area()) // Output: Area: 50
}
