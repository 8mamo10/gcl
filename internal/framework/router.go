package framework

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/8mamo10/gcl/internal/infra"
	"github.com/labstack/echo/v4"
)

const port = "8080"

type Server struct {
	Port string
	Echo *echo.Echo
}

func CreateServer() *Server {
	e := echo.New()
	e.GET("/users", getUsersHandler)
	e.GET("/accounts", getAccountsHandler)
	return &Server{Port: port, Echo: e}
}

func (srv *Server) Run() error {
	go func() {
		if err := srv.Echo.Start(":" + srv.Port); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return srv.Echo.Shutdown(ctx)
}

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

func getUsersHandler(ctx echo.Context) error {
	rows, err := infra.GetDB().Model(&User{}).Rows()
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			fmt.Sprintf("failed to query users: %v", err))
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var user User
		err := infra.GetDB().ScanRows(rows, &user)
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
	rows, err := infra.GetDB().Model(&Account{}).Rows()
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			fmt.Sprintf("failed to query accounts: %v", err))
	}
	defer rows.Close()

	var accounts []*Account
	for rows.Next() {
		var account Account
		err := infra.GetDB().ScanRows(rows, &account)
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
