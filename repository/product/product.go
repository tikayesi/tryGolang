package product

import (
	"database/sql"
	"log"

	guuid "github.com/google/uuid"
)

type Product struct {
	ProductId          string
	ProductName        string
	ProductDescription string
}

func FindProductByCode(db *sql.DB, code string) (*Product, error) {
	row := db.QueryRow(`SELECT p.id, p.product_name, p.product_description
	FROM product p WHERE p.id = ?`, code)

	p := new(Product)
	err := row.Scan(&p.ProductId, &p.ProductName, &p.ProductDescription)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	return p, nil
}

func CreateProduct(db *sql.DB, product Product) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}
	stmt, err := db.Prepare("INSERT INTO product(id,product_name,product_description)  VALUES(?, ?, ?)")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	defer stmt.Close()

	id := guuid.New()
	if _, err := stmt.Exec(id, product.ProductName, product.ProductDescription); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	return tx.Commit()
}
