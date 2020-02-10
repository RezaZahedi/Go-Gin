package product

type ProductDTO struct {
	ID          *uint  `json:"id,string"`
	Name        string `json:"name"`
	Description string `json:"description,string"`
}
