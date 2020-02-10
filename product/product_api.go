package product

import (
	"strconv"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ProductAPI struct {
	ProductService ProductService
}

func ProvideProductAPI(p ProductService) ProductAPI {
	return ProductAPI{ProductService: p}
}

// TODO: write all the handlers form template's reference

func (p *ProductAPI) FindAll(c *gin.Context) {
	products, err := p.ProductService.FindAll()
	if err != nil {
		c.Error(err)

	}
	c.JSON(http.StatusOK, gin.H{"products": ToProductDTOs(products)})
}

func (p *ProductAPI) FindByID(c *gin.Context) {
	id, _ :=  strconv.Atoi(c.Param("id"))
	product, _:= p.ProductService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"product": ToProductDTO(product)})
}

func (p *ProductAPI) Create(c *gin.Context) {
	var productDTO ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	product, err := ToProduct(productDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	createdProduct, _ := p.ProductService.Create(product)

	c.JSON(http.StatusOK, gin.H{"product": ToProductDTO(createdProduct)})
}

func (p *ProductAPI) Update(c *gin.Context) {
	var productDTO ProductDTO
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
