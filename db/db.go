package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDataBase() *sql.DB {
	connection := "root:Milkado@1@/rosharan_books"
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
