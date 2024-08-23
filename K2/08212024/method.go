package main

import (
	"errors"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// cùng kiểu có thể khai báo như dưới
func Plus(number1, number2 float64) float64 {
	return number1 + number2
}

func Divison(number1, number2 float64) (float64, error) {
	if number2 == 0 {
		return 0, errors.New("không có chia cho 0")
	}
	return number1 / number2, nil
}

type User struct {
	Name  string
	Email string
}

// reciever method
// method này thuộc đối tượng user
func (u User) GetName() string {
	return u.Name
}

func (u User) SetName(newName string) {
	u.Name = newName
}
func (u *User) SetNamePointer(newName string) {
	u.Name = newName
}

func main() {
	var res1 = Plus(9, 10)
	var res2, err = Divison(9, 3)
	if err != nil {
		spew.Dump(err)
		return
	}
	// var res3, err2 = Divison(9, 3)
	spew.Dump(res1)
	spew.Dump(res2)
	// spew.Dump(err)
	// spew.Dump(res3)
	// spew.Dump(err2)

	var user = User{
		Name:  "Manh",
		Email: "NguyenCongManh@gmail.com",
	}
	// fmt.Print(user)
	fmt.Printf("Name = %s \n", user.GetName())

	// chỉ thao tác trên bản copy của thằng user
	user.SetName("Nguyen")
	fmt.Printf("Name = %s \n", user.GetName())

	//pointer
	user.SetNamePointer("Nguyen")
	fmt.Printf("Name = %s \n", user.GetName())
}
