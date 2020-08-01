package product

import (
	"database/sql"
)

type IProductService interface {
	GetProductByCode(code string) *Product
}

type ProductService struct {
	db *sql.DB
}

func (ps *ProductService) GetProductByCode(code string) *Product {
	var product *Product
	var err error
	product, err = FindProductByCode(ps.db, code)

	if err != nil {
		return nil
	}
	return product
}

func NewProductService(db *sql.DB) IProductService {
	return &ProductService{db}
}
