package usecase

import (
	"errors"
	"log"

	"teste-go/internal/entity"
	"teste-go/internal/repository"
)

type UserUseCase interface {
	GetAll() ([]*entity.User, error)
	GetByID(id int) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id int) error
	ValidateCredentials(username, password string) (*entity.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

// ValidateCredentials validates the user's credentials
func (uc *userUseCase) ValidateCredentials(email, password string) (*entity.User, error) {
	user, err := uc.userRepository.GetByUsername(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Compare the provided password with the user's password
	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (uc *userUseCase) GetAll() ([]*entity.User, error) {
	users, err := uc.userRepository.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

func (uc *userUseCase) GetByID(id int) (*entity.User, error) {
	user, err := uc.userRepository.GetByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

func (uc *userUseCase) Create(user *entity.User) error {
	err := uc.userRepository.Create(user)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (uc *userUseCase) Update(user *entity.User) error {
	err := uc.userRepository.Update(user)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (uc *userUseCase) Delete(id int) error {
	err := uc.userRepository.Delete(id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
