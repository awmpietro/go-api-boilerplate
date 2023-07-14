package repository

import (
	"log"

	"gorm.io/gorm"

	"teste-go/internal/entity"
)

type ProductRepository interface {
	GetAll() ([]*entity.Product, error)
	GetByID(id int) (*entity.Product, error)
	Create(product *entity.Product) error
	Update(product *entity.Product) error
	Delete(id int) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (pr *productRepository) GetAll() ([]*entity.Product, error) {
	var products []*entity.Product
	result := pr.db.Find(&products)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return products, nil
}

func (pr *productRepository) GetByID(id int) (*entity.Product, error) {
	var product entity.Product
	result := pr.db.First(&product, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Println(result.Error)
		return nil, result.Error
	}

	return &product, nil
}

func (pr *productRepository) Create(product *entity.Product) error {
	result := pr.db.Create(product)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}

func (pr *productRepository) Update(product *entity.Product) error {
	result := pr.db.Save(product)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}

func (pr *productRepository) Delete(id int) error {
	result := pr.db.Delete(&entity.Product{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}
