package cmd

import (
	"github.com/Pkittipat/polynomial-service/config"
	"github.com/Pkittipat/polynomial-service/internal/http/handlers"
	"github.com/Pkittipat/polynomial-service/internal/usecases"
	"go.uber.org/dig"
)

// BuildContainer ...
func BuildContainer() *dig.Container {
	container := dig.New()

	_ = container.Provide(config.LoadConfig)

	// Handlers
	_ = container.Provide(handlers.NewPolynomialHandler)

	// Usecase
	_ = container.Provide(usecases.NewPolynomialUsecase)

	return container
}
