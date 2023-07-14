package entity

import "time"

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"not null"`
	Email     *string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
