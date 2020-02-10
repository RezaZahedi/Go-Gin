package REST_api

import (
	"github.com/RezaZahedi/Go-Gin/product"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type ProductAPI struct {
	ProductService product.ProductService
}

func ProvideProductAPI(p product.ProductService) ProductAPI {
	return ProductAPI{ProductService: p}
}

// TODO: write all the handlers form template's reference

func (p *ProductAPI) FindAll(c *gin.Context) {
	products, err := p.ProductService.FindAll()
	if err != nil {
		c.Error(err)

	}
	c.JSON(http.StatusOK, gin.H{"products": product.ToProductDTOs(products)})
}

func (p *ProductAPI) FindByID(c *gin.Context) {
	id, _ :=  strconv.Atoi(c.Param("id"))
	_product, _:= p.ProductService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"product": product.ToProductDTO(_product)})
}

func (p *ProductAPI) Create(c *gin.Context) {
	var productDTO product.ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	_product, err := product.ToProduct(productDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	createdProduct, _ := p.ProductService.Create(_product)

	c.JSON(http.StatusOK, gin.H{"product": product.ToProductDTO(createdProduct)})
}

func (p *ProductAPI) Update(c *gin.Context) {
	var productDTO product.ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ :=  strconv.Atoi(c.Param("id"))
	product, err := p.ProductService.FindByID(uint(id))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	product.Name = productDTO.Name
	product.Description = productDTO.Description
	p.ProductService.Update(uint(id), product)

	c.Status(http.StatusOK)
}

func (p *ProductAPI) Delete(c *gin.Context) {
	id, _ :=  strconv.Atoi(c.Param("id"))
	product, err := p.ProductService.FindByID(uint(id))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	p.ProductService.Delete(product)

	c.Status(http.StatusOK)
}
