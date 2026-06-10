package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/felipematheus1337/go-ecommerce-ms/internal/dto/request"
	"github.com/felipematheus1337/go-ecommerce-ms/internal/model"
	"github.com/felipematheus1337/go-ecommerce-ms/internal/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (ph *ProductHandler) CreateProduct(ctx *gin.Context) {

	var productRequest request.ProductRequest

	if err := ctx.ShouldBind(&productRequest); err != nil {
		ph.sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	product := model.Product{
		Name:      productRequest.Name,
		CreatedAt: time.Now(),
		Stock:     productRequest.Stock,
		Price:     productRequest.Price,
	}

	response, err := ph.service.Create(product)

	if err != nil {
		ph.sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ph.sendSuccess(ctx, "create-product", response, http.StatusCreated)

}

func (ph *ProductHandler) GetProducts(ctx *gin.Context) {
	res, err := ph.service.GetProducts()

	if err != nil {
		ph.sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ph.sendSuccess(ctx, "get-products", res, http.StatusOK)
}

func (ph *ProductHandler) sendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   message,
		"errorCode": code,
	})
}

func (ph *ProductHandler) sendSuccess(ctx *gin.Context, op string, data interface{}, code int) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message": fmt.Sprintf("operation from handler : %s successfull", op),
		"data":    data,
	})
}
