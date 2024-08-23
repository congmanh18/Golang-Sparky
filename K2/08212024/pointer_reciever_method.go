package main

import "fmt"

// Khai báo một struct
type Rectangle struct {
	width, height int
}

// Phương thức với receiver theo con trỏ
func (r *Rectangle) scale(factor int) {
	r.width *= factor
	r.height *= factor
}

func main() {
	rect := Rectangle{width: 10, height: 5}
	rect.scale(3)
	fmt.Println("Scaled Width:", rect.width)   // Output: Scaled Width: 20
	fmt.Println("Scaled Height:", rect.height) // Output: Scaled Height: 10
}
