package usecase

import (
	"log"

	"teste-go/internal/entity"
	"teste-go/internal/repository"
)

type ProductUseCase interface {
	GetAll() ([]*entity.Product, error)
	GetByID(id int) (*entity.Product, error)
	Create(product *entity.Product) error
	Update(product *entity.Product) error
	Delete(id int) error
}

type productUseCase struct {
	productRepository repository.ProductRepository
}

func NewProductUseCase(productRepository repository.ProductRepository) ProductUseCase {
	return &productUseCase{
		productRepository: productRepository,
	}
}

func (uc *productUseCase) GetAll() ([]*entity.Product, error) {
	products, err := uc.productRepository.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return products, nil
}

func (uc *productUseCase) GetByID(id int) (*entity.Product, error) {
	product, err := uc.productRepository.GetByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return product, nil
}

func (uc *productUseCase) Create(product *entity.Product) error {
	err := uc.productRepository.Create(product)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (uc *productUseCase) Update(product *entity.Product) error {
	err := uc.productRepository.Update(product)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (uc *productUseCase) Delete(id int) error {
	err := uc.productRepository.Delete(id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
