package db

import "database/sql"

func GetDB() *sql.DB {
	db, err := sql.Open("mysql", "u1107404_efishery:efishery123!@tcp(5.181.216.74:3306)/u1107404_efishery")

	if err != nil {
		panic(err.Error())
	}

	return db
}

func CloseDB(db *sql.DB) {
	db.Close()
}
