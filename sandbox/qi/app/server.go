package main

import "github.com/8mamo10/gcl/sandbox/qi/app/infrastructure"

func main() {
	infrastructure.Router.Run(":8088")
}
