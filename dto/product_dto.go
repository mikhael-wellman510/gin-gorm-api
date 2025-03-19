package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type (
	ProductResponse struct {
		Id                 string          `json:"id"`
		Name               string          `json:"name"`
		Price              decimal.Decimal `json:"price"`
		Brand              string          `json:"brand"`
		WeightProduct      decimal.Decimal `json:"weight_product"`
		StockProduct       uint8           `json:"stock_product"`
		DescriptionProduct string          `json:"description_product"`
		CreatedAt          time.Time       `json:"created_at"`
		UpdatedAt          time.Time       `json:"updated_at"`
	}

	ProductResponsePaggingAndFilter struct {
		ProductResponse []ProductResponse
		PaggingResponse PagginationResponse
	}

	/*
		binding : wajib di isi , jika tidak akan error
	*/
	ProductRequest struct {
		Id                 string          `json:"id"`
		Name               string          `json:"name" binding:"required"`
		Brand              string          `json:"brand" binding:"required"`
		Price              decimal.Decimal `json:"price"`
		WeightProduct      decimal.Decimal `json:"weight_product"`
		StockProduct       uint8           `json:"stock_product"`
		DescriptionProduct string          `json:"description_product"`
	}
)
