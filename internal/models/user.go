package models

import "github.com/google/uuid"

type User struct {
	ID uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`

	Email    string `gorm:"unique"`
	Password string
}
