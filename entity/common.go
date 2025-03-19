package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
MySql tidak mengenal tipe data uuid , jadi harus menggunakan varchar
*/
type Base struct {
	ID        string    `json:"id" gorm:"type:varchar(36);primary_key"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

/*
ini hook untuk generate UUID sebelum insert otomatis
*/
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {

	b.ID = uuid.New().String()
	fmt.Println("GENERATE UUID:", b.ID)
	return nil
}
