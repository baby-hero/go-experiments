package models

import (
	"fmt"
	"net/mail"
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	FullName  string         `gorm:"type:varchar(100);not null" json:"full_name"`
	Email     string         `gorm:"type:varchar(100);unique;not null" json:"email"`
	Phone     string         `gorm:"type:varchar(20)" json:"phone,omitempty"`
	Address   string         `gorm:"type:varchar(255)" json:"address,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (c *Customer) Validate() error {
	if c.FullName == "" {
		return fmt.Errorf("full name is required")
	}
	if _, err := mail.ParseAddress(c.Email); err != nil {
		return fmt.Errorf("invalid email format")
	}
	return nil
}
