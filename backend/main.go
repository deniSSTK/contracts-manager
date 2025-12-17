package main

import (
	"contracts-manager/internal/delivery/http"
	"contracts-manager/internal/infrastructure"
	"contracts-manager/internal/usecases"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		http.Module,
		infrastructure.Module,
		usecases.Module,
	)

	app.Run()
}
