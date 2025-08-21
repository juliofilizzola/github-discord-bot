package model

type NotificationStatus string

const (
	StatusPending NotificationStatus = "pending"
	StatusSent    NotificationStatus = "sent"
	StatusFailed  NotificationStatus = "failed"
)

type Notification struct {
	BaseModel
	Message string             `gorm:"type:text;not null"`
	Status  NotificationStatus `gorm:"type:enum('pending','sent','failed');default:'pending';not null"`
	Type    string             `gorm:"type:varchar(50);not null"`
	UserID  uint               `gorm:"not null"`
}

func init() {
	println("Registering Notification model Notification")
	RegisterModel(&Notification{})
}
