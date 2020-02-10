package product

import (
	"errors"
	"github.com/RezaZahedi/Go-Gin/model"
)

func ToProduct(productDTO ProductDTO) (Product, error) {
	if productDTO.ID == nil {
		return nil, errors.New("ID is required")
	}
	return Product(&model.Book{
		ID:          model.ID{BackField: int(*productDTO.ID)},
		Name:        productDTO.Name,
		Description: productDTO.Description,
	}), nil
}

func ToProductDTO(product Product) ProductDTO {
	id := uint(product.ID.BackField)
	return ProductDTO{
		ID:          &id,
		Name:        product.Name,
		Description: product.Description,
	}
}

func ToProductDTOs(products []Product) []ProductDTO {
	productdtos := make([]ProductDTO, len(products))

	for i, itm := range products {
		productdtos[i] = ToProductDTO(itm)
	}

	return productdtos
}
