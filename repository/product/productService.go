package product

type IProductService interface {
	GetProductByCode(code string) *Product
}

type ProductService struct {
	productRepo IProductRepository
}

func (ps *ProductService) GetProductByCode(code string) *Product {
	var product *Product
	var err error
	product, err = ps.productRepo.FindProductByCode(code)

	if err != nil {
		return nil
	}
	return product
}

func NewProductService(productRepo IProductRepository) IProductService {
	return &ProductService{productRepo}
}
