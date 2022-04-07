package model

type UserTw struct {
	TwId   int64 `gorm:"primaryKey"`
	UserId uint  `gorm:"index"`
}
