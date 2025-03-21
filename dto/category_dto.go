package dto

import "time"

type (
	CategoryResponse struct {
		Id           string    `json:"id"`
		CategoryName string    `json:"category_name"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	CategoryRequest struct {
		CategoryName string `json:"category_name" binding:"required"`
	}
)
