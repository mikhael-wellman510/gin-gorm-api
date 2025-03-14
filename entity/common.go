package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

/*
ini hook untuk generate UUID sebelum insert otomatis
*/
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {

	b.ID = uuid.New().String()

	return nil
}
