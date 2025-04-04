package repository

import (
	"GinGonicGorm/entity"
	"context"
	"fmt"
	"math"

	"gorm.io/gorm"
)

type (
	ProductRepository interface {
		SaveProduct(ctx context.Context, tx *gorm.DB, product entity.Product) (entity.Product, error)
		FindProductById(ctx context.Context, productId string) (entity.Product, error)
		FindAllProduct(ctx context.Context) ([]entity.Product, error)
		UpdatedProduct(ctx context.Context, tx *gorm.DB, product entity.Product) (entity.Product, error)
		PagginationAndSearchProduct(ctx context.Context, name string, limit int, offset int) ([]entity.Product, int, error)
	}

	// Attribute
	productRepository struct {
		db *gorm.DB
	}
)

func NewProductRepository(db *gorm.DB) ProductRepository {

	return &productRepository{
		db: db,
	}

}

/*
jadi fungsi tx yaitu untuk membatalkan transaksi jika browser d tutup /
atau jika postman d cancel karena karena kelamaan get data
*/
func (r *productRepository) SaveProduct(ctx context.Context, tx *gorm.DB, product entity.Product) (entity.Product, error) {

	if tx == nil {

		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&product).Error; err != nil {
		fmt.Println("Masuk ke sini kerena err :", err)
		return entity.Product{}, err
	}

	return product, nil
}

func (r *productRepository) FindProductById(ctx context.Context, productId string) (entity.Product, error) {

	// Define dulu product nta
	product := entity.Product{}
	// ini untuk cari primary key kalau pake uuid
	if err := r.db.WithContext(ctx).Preload("Category").First(&product, "id = ?", productId).Error; err != nil {

		return entity.Product{}, err
	}

	return product, nil
}

func (r *productRepository) FindAllProduct(ctx context.Context) ([]entity.Product, error) {

	var product []entity.Product

	// Preload untuk ambil join table nya
	if err := r.db.WithContext(ctx).Preload("Category").Find(&product).Error; err != nil {

		return nil, err
	}

	return product, nil
}

func (r *productRepository) UpdatedProduct(ctx context.Context, tx *gorm.DB, product entity.Product) (entity.Product, error) {

	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&product).Error; err != nil {
		return entity.Product{}, err
	}
	return product, nil

}

func (r *productRepository) PagginationAndSearchProduct(ctx context.Context, name string, size int, page int) ([]entity.Product, int, error) {

	var product []entity.Product
	var totalItems int64

	// Menggunakan preload karena join table , asli nya ga pake
	query := r.db.WithContext(ctx).Preload("Category").Model(&entity.Product{})

	if name != "" {
		query.Where("name LIKE ?", "%"+name+"%")
	}

	if page <= 0 {
		page = 1
	}
	if err := query.Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}
	totalPage := int(math.Ceil(float64(totalItems) / float64(size)))

	offset := (page - 1) * size

	query = query.Limit(size).Offset(offset)

	err := query.Order("name ASC").Find(&product).Error

	if page > totalPage {
		page = 1
	}
	if err != nil {

		return []entity.Product{}, 0, err
	}

	return product, int(totalItems), nil
}
