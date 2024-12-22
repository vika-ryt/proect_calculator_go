package main

import (
	"github.com/vika-ryt/proect_calculator_go/internal/application"
)

func main() {
	app := application.New()
	// app.Run()
	app.RunServer()
}