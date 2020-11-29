package db

import (
	"database/sql"
	"os"
)

func GetDB() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DB_CONFIG"))

	if err != nil {
		panic(err.Error())
	}

	return db
}

func CloseDB(db *sql.DB) {
	db.Close()
}
