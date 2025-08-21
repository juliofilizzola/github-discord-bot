package model

import "github.com/juliofilizzola/github-discord-bot/internal/constants"

type Notification struct {
	BaseModel
	Message string                       `gorm:"type:text;not null"`
	Status  constants.NotificationStatus `gorm:"type:enum('pending','sent','failed');default:'pending';not null"`
	Type    string                       `gorm:"type:varchar(50);not null"`
	UserID  uint                         `gorm:"not null" json:"user_id" gorm:"column:user_id;foreignKey:UserID;references:ID" ref:"githubUsers"`
}

func init() {
	println("Registering Notification model Notification")
	RegisterModel(&Notification{})
}
