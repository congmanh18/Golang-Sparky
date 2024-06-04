// public database intance cua gorm

package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Name string
}

// 2 phong cach
func NewDB() (*gorm.DB, error) {
	gormDb, err := gorm.Open(postgres.Open("dsn"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connect database: %v, err")
	}

	return gormDb, nil
	// var user = User{
	// 	Name: "Manh",
	// }
	// gormDb.Create(&user)

}
