package main

import (
	"database/sql"
	"fmt"
	"log"
	"vpgosql/config"
	"vpgosql/repository"
	"vpgosql/repository/product"
)

type vpGoSql struct {
	db *sql.DB
}

//constructor
func NewVpGoSql(c *config.Conf) *vpGoSql {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Db.DbUser, c.Db.DbPassword, c.Db.DbHost, c.Db.DbPort, c.Db.SchemaName)
	db, err := repository.NewDbInitialization(c.Db.DbEngine, connString).InitDB()
	if err != nil {
		log.Panic(err)
	}
	return &vpGoSql{
		db,
	}
}

//receiver
func (v *vpGoSql) run() {
	fmt.Println("Go Sql")

	prodService := product.NewProductService(v.db)
	product := prodService.GetProductByCode("1")
	fmt.Printf("Product: %s %s", product.ProductId, product.ProductName)
}

func main() {
	conf := config.NewAppConfig()
	NewVpGoSql(conf).run()
}
