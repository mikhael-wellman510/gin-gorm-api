package entity

import (
	"github.com/shopspring/decimal"
)

type Product struct {
	Name               string          `json:"name" gorm:"column:name;not null;unique"`
	Brand              string          `json:"brand" gorm:"column:brand"`
	Price              decimal.Decimal `json:"price" gorm:"column:price;type:decimal(15,2);not null"`
	WeightProduct      decimal.Decimal `json:"weight_product" gorm:"column:weight_product;type:decimal(10,2)"`
	StockProduct       uint8           `json:"stock_product" gorm:"column:stock_product;default:0"`
	DescriptionProduct string          `json:"description_product" gorm:"column:description_product;type:text"`
	Base
}
