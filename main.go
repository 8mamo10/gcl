package main

import (
	"os"

	"github.com/8mamo10/gcl/internal/framework"
	"github.com/8mamo10/gcl/internal/infra"
)

func init() {
	infra.InitializeDatabase()
}

func main() {
	srv := framework.CreateServer()
	if err := srv.Run(); err != nil {
		os.Exit(1)
	}
}
