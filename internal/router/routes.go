package router

import (
	"github.com/felipematheus1337/go-ecommerce-ms/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, ph *handler.ProductHandler) {

	basePath := "/api/v1/products"

	v1 := router.Group(basePath)

	RegisterProductsRoutes(v1, ph)
}

func RegisterProductsRoutes(v1 *gin.RouterGroup, ph *handler.ProductHandler) {
	{
		v1.POST("/", ph.CreateProduct)
		v1.GET("/", ph.GetProducts)
	}
}
