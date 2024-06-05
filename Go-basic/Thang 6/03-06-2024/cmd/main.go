package main

import (
	"fmt"
	"microlap-gorm/db"
	"microlap-gorm/entity"
	"microlap-gorm/store"
)

var enableMigration = false

func main() {
	var conn = db.Connection{
		Host:     "localhost",
		User:     "spadyprod",
		Password: "123456",
		DBName:   "gorm_db",
		Port:     "5432",
	}

	var gorm, err = db.NewDB(conn)
	if err != nil {
		panic(err)
	}

	if enableMigration {
		gorm.Debug().AutoMigrate(&entity.User{})
	}

	var userStore = store.NewUserStore(gorm)

	// test insert user
	var user = entity.User{
		ID:    "678",
		Name:  "Ryan",
		Email: "ryan@gmail.com",
		Phone: "0973901735",
	}
	if err := userStore.Save(user); err != nil {
		fmt.Printf("error insert record: %v \n", err)
	} else {
		fmt.Printf("record inserted: %v \n", user)
	}

	// test update
	var updateUser = entity.User{
		//ID:    "345",
		Name:  "Ryan 3456",
		Email: "ryan@gmail.com",
		Phone: "0973901735",
	}
	if err := userStore.UpdateV2("345", updateUser); err != nil {
		fmt.Printf("error update record: %v \n", err)
	} else {
		fmt.Printf("record updated: %v \n", user)
	}

}
