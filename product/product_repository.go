package product

import (
	"github.com/RezaZahedi/Go-Gin/database"
	"github.com/RezaZahedi/Go-Gin/model"
)

type ProductRepository struct {
	BookDB *database.BookDB
}

func ProvideProductRepository(db *database.BookDB) ProductRepository {
	return ProductRepository{BookDB: db}
}

func (p *ProductRepository) FindAll() ([]Product, error) {
	books, err := p.BookDB.FindAll()
	products := make([]Product, 0, len(books))
	for _, book := range books {
		products = append(products, Product(book))
	}
	return products, err
}

func (p *ProductRepository) FindByID(id uint) (Product, error) {
	return p.BookDB.FindByID(&model.ID{BackField: int(id)})
}

func (p *ProductRepository) Create(id uint, product Product) (Product, error) {
	return p.BookDB.Create(&model.ID{BackField: int(id)}, product)
}

func (p *ProductRepository) Update(id uint, product Product) (Product, error) {
	return p.BookDB.Update(&model.ID{BackField: int(id)}, product)
}

func (p *ProductRepository) Delete(product Product) error {
	return p.BookDB.Delete(&product.ID)
}
