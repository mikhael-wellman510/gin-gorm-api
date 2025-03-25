package entity

type Photo struct {
	Base
	FileName string `json:"filename"`
	Url      string `json:"url"`
}
