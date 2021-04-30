// models/task.go

package models

import (
	"time"
)

type Engineer struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	UUID        string    `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Skill       string    `json:"skill"`
	Address     string    `json:"address"`
	DateOfBirth string    `json:"dateofbirth"`
	PhoneNumber string    `json:"phonenumber"`
	Showcase    string    `json:"showcase"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
