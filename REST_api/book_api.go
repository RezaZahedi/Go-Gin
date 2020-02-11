package REST_api

import (
	"github.com/RezaZahedi/Go-Gin/product"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductAPI struct {
	ProductService product.ProductService
}

func ProvideProductAPI(p product.ProductService) *ProductAPI {
	return &ProductAPI{ProductService: p}
}

func (p *ProductAPI) ShowIndexPage(c *gin.Context)  {
	books, err := p.ProductService.FindAll()
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	// Call the render function with the name of the template to render
	render(c,
		gin.H{"title": "Home Page",
			"payload": product.ToProductDTOs(books)},
		"index.html")
}

func (*ProductAPI) ShowBookCreatingPage(c *gin.Context) {
	render(c,
		gin.H{"title": "Create New Book"},
		"create-book.html")
}

func (p *ProductAPI) GetBook(c *gin.Context) {
	bookId, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	book, err := p.ProductService.FindByID(uint(bookId))
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	bookDTO := product.ToProductDTO(book)
	render(c,
		gin.H{"title": bookDTO.Name,
			"payload": bookDTO},
		"book.html")

}

func (p *ProductAPI) CreateBook(c *gin.Context) {
	CreateOrUpdateBook(c, p.ProductService.Create)
}

func (p *ProductAPI) UpdateBook(c *gin.Context) {
	CreateOrUpdateBook(c, p.ProductService.Update)
}

func CreateOrUpdateBook(c *gin.Context, tempAction func(id uint, product product.Product) (product.Product, error)) {
	name := c.PostForm("name")
	description := c.PostForm("description")
	bookID, err := strconv.Atoi(c.PostForm("book_id"))
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	bookIDUint := uint(bookID)
	productDTO := product.ProductDTO{
		ID:          &bookIDUint,
		Name:        name,
		Description: description,
	}
	_product, err := product.ToProduct(productDTO)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	_, err = tempAction(bookIDUint, _product)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	render(c,
		gin.H{"title": "Submission Successful",
			"payload": productDTO},
		"submission-successful.html")
}

func (p *ProductAPI) DeleteBook(c *gin.Context) {
	bookId, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	book, err := p.ProductService.FindByID(uint(bookId))
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	err = p.ProductService.Delete(book)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	render(c,
		gin.H{"title": "Deletion Successful"},
		"deletion_successful.html")
}

