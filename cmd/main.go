package main

import (
	"github.com/rabboni171/transactions-demo/pkg/di"
)

func main() {
	app := di.InitializeApp()
	app.Run()

}
