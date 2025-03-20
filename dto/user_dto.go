package dto

import "time"

type (
	UserResponse struct {
		Id          string    `json:"id"`
		Username    string    `json:"username"`
		Email       string    `json:"email"`
		Password    string    `json:"password"`
		PhoneNumber string    `json:"phone_number"`
		IsActice    bool      `json:"is_active"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatetAt   time.Time `json:"updated_at"`
	}

	UserRequest struct {
		Username    string `json:"username" binding:"required"`
		Email       string `json:"email" binding:"required,email"`
		Password    string `json:"password" binding:"required"`
		PhoneNumber string `json:"phone_number" binding:"required"`
	}
)
