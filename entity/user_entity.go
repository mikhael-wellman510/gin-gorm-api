package entity

type User struct {
	Username    string `json:"username" gorm:"column:username;not null;unique"`
	Email       string `json:"email" gorm:"column:email;not null;unique"`
	Password    string `json:"password" gorm:"column:password;not null"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number;not null"`
	IsActive    bool   `json:"is_active" gorm:"column:is_active;default:false"`
	Base
}
