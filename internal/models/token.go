package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Token struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`

	UserID uuid.UUID `gorm:"type:uuid"`
	UserIP string

	AccessToken  string
	RefreshToken string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Token) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()

	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	tx.Save(t)

	return
}
