package param

import "time"

type UserInfo struct {
	ID          uint      `json:"id"`
	PhoneNumber string    `json:"phone_number"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
}
