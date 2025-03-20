package repository

import (
	"GinGonicGorm/entity"
	"context"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(ctx context.Context, user entity.User) (entity.User, error)
		FindByEmail(ctx context.Context, email string) (entity.User, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {

	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(ctx context.Context, user entity.User) (entity.User, error) {

	if err := ur.db.WithContext(ctx).Create(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (ur *userRepository) FindByEmail(ctx context.Context, email string) (entity.User, error) {

	var user entity.User

	if err := ur.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {

		return entity.User{}, err
	}

	return user, nil

}
