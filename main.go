package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error
	_, err = sql.Open("mysql", "root:password@tcp(mysql:3307)/interview_test")
	if err != nil {
		panic(err)
	}
	println("Hello")
}
