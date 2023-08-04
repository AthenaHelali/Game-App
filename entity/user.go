package entity

import "time"

type User struct {
	ID          uint
	PhoneNumber string
	Name        string
	CreatedAt   time.Time
}
