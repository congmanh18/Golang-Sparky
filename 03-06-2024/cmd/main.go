package main

import "microlap-gorm/db"

func main() {
	var conn = db.Connection{
		Host:     "localhost",
		User:     "postgres",
		Password: "231002",
		DBName:   "postgres",
		Port:     "5432",
	}

	var gorm, err = db.NewDB(conn)

	if err != nil {
		panic(err)
	}

	gorm.Debug().AutoMigrate(&entity.User{})
}
