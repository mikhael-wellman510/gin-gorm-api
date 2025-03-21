package entity

type Category struct {
	Base
	CategoryName string `json:"category_name" gorm:"not null;unique"`
}
