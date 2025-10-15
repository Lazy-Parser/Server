package entity

import "time"

type User struct {
	ID               uint      `gorm:"primarykey"`
	Username         string    `gorm:"column:username"`
	Password         string    `gorm:"column:password"`           // remember that its a one-time password
	LastPasswordTime time.Time `gorm:"column:last_password_time"` // time, when the last password was created. Needed to make some ttl for the password
	RoleID           uint
	Role             Role `gorm:"foreignKey:RoleID"`
}
