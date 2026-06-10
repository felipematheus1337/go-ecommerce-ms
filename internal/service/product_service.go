package service

import (
	"github.com/felipematheus1337/go-ecommerce-ms/internal/dto/response"
	"github.com/felipematheus1337/go-ecommerce-ms/internal/model"
	"github.com/felipematheus1337/go-ecommerce-ms/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func (s ProductService) Create(product model.Product) (response.ProductResponse, error) {
	var productResponse *response.ProductResponse

	if err := s.repo.Create(&product).Error; err != nil {
		return response.ProductResponse{}, nil
	}

	productResponse = &response.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Stock:     product.Stock,
		CreatedAt: product.CreatedAt,
		Price:     product.Price,
	}

	return *productResponse, nil

}

func (s ProductService) GetProducts() ([]response.ProductResponse, error) {
	var products []model.Product

	if err := s.repo.GetAll(&products).Error; err != nil {
		return nil, err
	}

	var listProductResponse []response.ProductResponse

	for _, emp := range products {
		listProductResponse = append(listProductResponse, response.ProductResponse{
			ID:        emp.ID,
			Name:      emp.Name,
			Price:     emp.Price,
			Stock:     emp.Stock,
			CreatedAt: emp.CreatedAt,
		})
	}

	return listProductResponse, nil

}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}
