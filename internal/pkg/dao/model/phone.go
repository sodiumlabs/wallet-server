package model

type UserPhone struct {
	Phone  string `gorm:"primaryKey;type:varchar(170)"`
	UserId uint   `gorm:"index"`
}
