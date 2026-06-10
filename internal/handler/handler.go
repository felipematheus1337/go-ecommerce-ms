package handler

import "github.com/felipematheus1337/go-ecommerce-ms/internal/service"

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}
