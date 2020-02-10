package main

import (
	"github.com/RezaZahedi/Go-Gin/database"
	"github.com/RezaZahedi/Go-Gin/product"
	"github.com/google/wire"
)

func initProductAPI(db *database.BookDB) product.ProductAPI {
	wire.Build(product.ProvideProductRepository, product.ProvideProductService, product.ProvideProductAPI)

	return product.ProductAPI{}
}
