package repository

import (
	"github.com/felipematheus1337/go-ecommerce-ms/internal/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func (r ProductRepository) GetAll(products *[]model.Product) *gorm.DB {

	return r.db.Find(products)
}

func (r ProductRepository) Create(m *model.Product) *gorm.DB {
	return r.db.Create(m)

}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}
