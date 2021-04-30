// models/task.go

package models

import (
	"time"
)

type Enterprice struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	UUID        string    `json:"uuid"`
	Name        string    `json:"name"`
	Logo        string    `json:"logo"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	PhoneNumber string    `json:"phonenumber"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
