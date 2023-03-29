package main

import (
	"fmt"
	"net/http"

	"github.com/brunacotrim/api-product/configs"
	"github.com/brunacotrim/api-product/internal/entity"
	"github.com/brunacotrim/api-product/internal/infra/database"
	"github.com/brunacotrim/api-product/internal/infra/database/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	fmt.Println(config)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
