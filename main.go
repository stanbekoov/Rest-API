package main

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CardOwner struct {
	ID   int
	Name string
	Age  int
}

func main() {
	dsn := "host=___ user=___ dbname=___ password=___ sslmode=disable"
	sqldb, err := sql.Open("pgx", dsn)
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{})

	if err != nil {
		panic("error")
	}

	db.Migrator().CreateTable(&CardOwner{})

	carl := &CardOwner{0, "Carl", 32}
	db.Create(carl)
}
