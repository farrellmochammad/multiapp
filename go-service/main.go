package main

import (
	deliver "go-service/delivery"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := deliver.SetupRouter()
	router.Run(":8013")
}
