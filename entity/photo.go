package entity

import "time"

type Photo struct {
	Base
	FileName string `json:"filename"`
	Url      string `json:"url"`
}

type PhotoResponse struct {
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
