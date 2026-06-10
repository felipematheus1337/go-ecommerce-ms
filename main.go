package main

import (
	"fmt"
	"log"
	"os"

	"github.com/felipematheus1337/go-ecommerce-ms/internal/config"
	"github.com/felipematheus1337/go-ecommerce-ms/internal/handler"
	"github.com/felipematheus1337/go-ecommerce-ms/internal/repository"
	"github.com/felipematheus1337/go-ecommerce-ms/internal/router"
	"github.com/felipematheus1337/go-ecommerce-ms/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {

	err := config.Init()

	if err != nil {
		fmt.Println(err)
		return
	}

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	db, err := config.InitializePostgres()

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	repo := repository.NewProductRepository(db)

	services := service.NewProductService(repo)

	handler := handler.NewProductHandler(services)

	router.InitializeRoutes(r, handler)

	err = r.Run(":" + port)

	if err != nil {
		fmt.Println(err)
		return
	}
}
