package service

import (
	"GinGonicGorm/dto"
	"GinGonicGorm/entity"
	"GinGonicGorm/repository"
)

type (
	CategoryService interface {
		SaveCategory(req dto.CategoryRequest) (dto.CategoryResponse, error)
		FindCategoryById(categoryId string) (entity.Category, error)
	}

	categoryService struct {
		categoryRepository repository.CategoryRepository
	}
)

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {

	return &categoryService{
		categoryRepository: categoryRepository,
	}
}

func (cs *categoryService) SaveCategory(req dto.CategoryRequest) (dto.CategoryResponse, error) {

	category := entity.Category{
		CategoryName: req.CategoryName,
	}

	res, err := cs.categoryRepository.SaveCategory(category)

	if err != nil {
		return dto.CategoryResponse{}, err
	}

	return dto.CategoryResponse{
		Id:           res.ID,
		CategoryName: res.CategoryName,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
	}, nil
}

func (cs *categoryService) FindCategoryById(categoryId string) (entity.Category, error) {

	res, err := cs.categoryRepository.FindCategoryById(categoryId)

	if err != nil {
		return entity.Category{}, err
	}

	return res, nil
}
