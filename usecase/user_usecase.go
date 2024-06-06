package usecase

import (
	"clean_arch/domain"

	"gorm.io/gorm"
)

type UserUseCase interface {
	GetUserByID(id uint) (domain.User, error)
}

type userUseCase struct {
	db *gorm.DB
}

func NewUserUseCase(db *gorm.DB) UserUseCase {
	return &userUseCase{
		db: db,
	}
}

func (uc *userUseCase) GetUserByID(id uint) (domain.User, error) {
	var user domain.User
	result := uc.db.First(&user, id)
	return user, result.Error
}
