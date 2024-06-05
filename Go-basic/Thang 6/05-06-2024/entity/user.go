package entity

type User struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Email string
	Phone string `gorm:"unique"`
}
