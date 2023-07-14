package repository

import (
	"errors"
	"log"

	"gorm.io/gorm"

	"teste-go/internal/entity"
)

type UserRepository interface {
	GetAll() ([]*entity.User, error)
	GetByID(id int) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id int) error
	GetByUsername(username string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) GetByUsername(email string) (*entity.User, error) {
	var user entity.User
	result := ur.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // User not found
		}
		return nil, result.Error // Other error occurred
	}

	return &user, nil
}

func (ur *userRepository) GetAll() ([]*entity.User, error) {
	var users []*entity.User
	result := ur.db.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return users, nil
}

func (ur *userRepository) GetByID(id int) (*entity.User, error) {
	var user entity.User
	result := ur.db.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Println(result.Error)
		return nil, result.Error
	}

	return &user, nil
}

func (ur *userRepository) Create(user *entity.User) error {
	result := ur.db.Create(user)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}

func (ur *userRepository) Update(user *entity.User) error {
	result := ur.db.Save(user)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}

func (ur *userRepository) Delete(id int) error {
	result := ur.db.Delete(&entity.User{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}
