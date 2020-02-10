package main

import (
	"github.com/RezaZahedi/Go-Gin/REST_api"
	"github.com/RezaZahedi/Go-Gin/database"
	"github.com/RezaZahedi/Go-Gin/product"
	"github.com/google/wire"
)

func initProductAPI(db *database.BookDB) REST_api.ProductAPI {
	wire.Build(product.ProvideProductRepository, product.ProvideProductService, REST_api.ProvideProductAPI)

	return REST_api.ProductAPI{}
}
