package product

type ProductService struct {
	ProductRepository ProductRepository
}

func ProvideProductService(p ProductRepository) ProductService {
	return ProductService{ProductRepository: p}
}

func (p *ProductService) FindAll() ([]Product, error) {
	return p.ProductRepository.FindAll()
}

func (p *ProductService) FindByID(id uint) (Product, error) {
	return p.ProductRepository.FindByID(id)
}

func (p *ProductService) Create(id uint, product Product) (Product, error) {
	return p.ProductRepository.Create(id, product)
}

func (p *ProductService) Update(id uint, product Product) (Product, error) {
	return p.ProductRepository.Update(id, product)
}

func (p *ProductService) Delete(product Product) error {
	return p.ProductRepository.Delete(product)
}
