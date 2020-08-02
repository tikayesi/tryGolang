package product

import (
	"database/sql"
	"log"
)

type IProductRepository interface {
	FindProductByCode(string) (*Product, error)
}

type ProductRepository struct {
	db *sql.DB
}

func (pr *ProductRepository) FindProductByCode(code string) (*Product, error) {
	row := pr.db.QueryRow(`
  SELECT p.id,p.product_name,p.product_description
  FROM product p WHERE p.id = ?`, code)

	p := new(Product)
	err := row.Scan(&p.ProductId, &p.ProductName, &p.ProductDescription)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	return p, nil
}

// func (pr *ProductRepository) CreateProduct(product Product) (*Product, error){

// }

func NewProductRepo(db *sql.DB) IProductRepository {
	return &ProductRepository{db}
}
