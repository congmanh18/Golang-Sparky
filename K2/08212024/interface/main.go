package main

import "fmt"

type UserRepo interface {
	GetUserByUID(userID string) string
	CreateUser(userID string) error
}

type UserRepoImpl struct {
}

func (u *UserRepoImpl) GetUserByUID(userID string) string {
	return userID
}

func (u *UserRepoImpl) CreateUser(userID string) error {
	return nil
}

func GetUser(repo UserRepo) string {
	return repo.GetUserByUID("100")
}

func main() {
	var userRepo = UserRepoImpl{}
	var user = GetUser(&userRepo)
	fmt.Println(user)
}

type ILogin interface {
	Excute(phone, password string) string
}

type User struct {
}

func (u *User) Excute(phone, password string) string {
	return fmt.Sprintf("Phone: %s, Password: %s", phone, password)
}

func Login(excute ILogin, phone, password string) string {
	return excute.Excute(phone, password)
}

func TestLogin() {
	var user = &User{}
	fmt.Println(Login(user, "123456789", "123456"))
}
