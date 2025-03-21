package repository

import (
	"GinGonicGorm/entity"

	"gorm.io/gorm"
)

type (
	CategoryRepository interface {
		SaveCategory(category entity.Category) (entity.Category, error)
		FindCategoryById(cateoryId string) (entity.Category, error)
	}

	categoryRepository struct {
		db *gorm.DB
	}
)

func NewCategoryRepository(db *gorm.DB) CategoryRepository {

	return &categoryRepository{
		db: db,
	}
}

func (cr *categoryRepository) SaveCategory(category entity.Category) (entity.Category, error) {

	if err := cr.db.Create(&category).Error; err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func (cr *categoryRepository) FindCategoryById(cateoryId string) (entity.Category, error) {

	categoryData := entity.Category{}

	if err := cr.db.First(&categoryData, "id = ?", cateoryId).Error; err != nil {
		return entity.Category{}, err
	}

	return categoryData, nil
}
