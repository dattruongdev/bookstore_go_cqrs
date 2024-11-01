package database

import (
	"log"

	"github.com/dattruongdev/bookstore_cqrs/config"
	"github.com/jmoiron/sqlx"
)

// datasource = user=foo dbname=bar sslmode=disable

func Connect(conf *config.Config) *sqlx.DB {
	log.Println(conf)
	connStr := conf.GetConnectionString()
	db, err := sqlx.Connect(conf.DriverName, connStr)

	if err != nil {
		log.Fatalln("Error occurred on sql connection. Check in database.NewConfig", err)
	}

	return db
}
