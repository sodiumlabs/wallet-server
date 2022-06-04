package model

type MPC struct {
	NodeId string `gorm:"primaryKey;type:varchar(170)"`
	Config string `gorm:"type:BLOB"`
}
