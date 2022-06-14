package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	ID   int64
	Name string
}

type User struct {
	ID        int64
	AccountID int64
	Name      string
	Email     string
}

var db *gorm.DB

func main() {
	dsn := "root:password@tcp(mysql:3306)/mrhc"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d

	e := echo.New()
	e.GET("/users", getUsersHandler)
	e.GET("/accounts", getAccountsHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

func getUsersHandler(ctx echo.Context) error {
	rows, err := db.Model(&User{}).Rows()
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			fmt.Sprintf("failed to query users: %v", err))
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var user User
		err := db.ScanRows(rows, &user)
		if err != nil {
			return echo.NewHTTPError(
				http.StatusBadRequest,
				fmt.Sprintf("failed retrieve data: %v", err))
		}
		users = append(users, &user)
	}

	ctx.JSON(http.StatusOK, users)
	return err
}

func getAccountsHandler(ctx echo.Context) error {
	rows, err := db.Model(&Account{}).Rows()
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			fmt.Sprintf("failed to query accounts: %v", err))
	}
	defer rows.Close()

	var accounts []*Account
	for rows.Next() {
		var account Account
		err := db.ScanRows(rows, &account)
		if err != nil {
			return echo.NewHTTPError(
				http.StatusBadRequest,
				fmt.Sprintf("failed retrieve data: %v", err))
		}
		accounts = append(accounts, &account)
	}

	ctx.JSON(http.StatusOK, accounts)
	return err
}
