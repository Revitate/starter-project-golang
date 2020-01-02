package model

type User struct {
	ID   string
	Name string `gorm:"name"`
}
