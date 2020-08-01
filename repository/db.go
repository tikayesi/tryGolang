package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type dbInitialization struct {
	dbEngine       string
	dataSourceName string
}

func NewDbInitialization(engine, dataSource string) *dbInitialization {
	return &dbInitialization{engine, dataSource}
}

func (dbi *dbInitialization) InitDB() (*sql.DB, error) {
	db, err := sql.Open(dbi.dbEngine, dbi.dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	//Ping = check database availability
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	return db, nil
}
