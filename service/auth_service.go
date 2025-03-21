package service

import (
	"GinGonicGorm/dto"
	"GinGonicGorm/entity"
	"GinGonicGorm/repository"
	"GinGonicGorm/utils"
	"context"
	"errors"
	"fmt"
)

type (
	AuthService interface {
		RegisterAccount(ctx context.Context, req dto.UserRequest) (dto.UserResponse, error)
		LoginAccount(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
	}

	authService struct {
		userRepository repository.UserRepository
	}
)

func NewAuthService(userRepository repository.UserRepository) AuthService {

	return &authService{
		userRepository: userRepository,
	}
}

func (as *authService) RegisterAccount(ctx context.Context, req dto.UserRequest) (dto.UserResponse, error) {

	// Hash password

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {

		return dto.UserResponse{}, err
	}

	fmt.Println("Err : ", err)
	userSave := entity.User{
		Username:    req.Username,
		Email:       req.Email,
		Password:    hashedPassword,
		PhoneNumber: req.PhoneNumber,
		IsActive:    true,
	}

	res, err := as.userRepository.Create(ctx, userSave)

	if err != nil {
		return dto.UserResponse{}, err
	}
	return dto.UserResponse{
		Id:          res.ID,
		Username:    res.Username,
		Email:       res.Email,
		Password:    res.Password,
		PhoneNumber: res.PhoneNumber,
		IsActice:    res.IsActive,
		CreatedAt:   res.CreatedAt,
		UpdatetAt:   res.UpdatedAt,
	}, nil
}

func (as *authService) LoginAccount(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {

	// CHECK EMAIL
	user, err := as.userRepository.FindByEmail(ctx, req.Email)

	if err != nil {
		return dto.LoginResponse{}, errors.New("INVALID EMAIL")
	}

	// Check Password

	checkPassword := utils.CheckPasswordHash(user.Password, req.Password)

	if !checkPassword {

		return dto.LoginResponse{}, errors.New("INVALID PASSWORD")
	}

	// Generate JWT TOKEN

	token, _ := utils.GenerateToken(user)

	return dto.LoginResponse{
		Token: token,
	}, nil
}
