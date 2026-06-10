package service

import "github.com/felipematheus1337/go-ecommerce-ms/internal/repository"

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}
