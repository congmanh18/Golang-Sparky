package repository

import "basic/user-service/entity"

type UserRepo struct {
}

func NewUserRepo() UserRepo {
	return UserRepo{}
}

func (u UserRepo) GetUserByUID(ID string) entity.User {
	return entity.User{
		ID:    &ID,
		Name:  "Manh",
		Email: "nguyenmanh@gmail.com",
	}
}
