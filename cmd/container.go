package cmd

import (
	"github.com/Pkittipat/polynomial-service/config"
	"github.com/Pkittipat/polynomial-service/internal/http/handlers"
	"github.com/Pkittipat/polynomial-service/internal/usecases"
	"github.com/Pkittipat/polynomial-service/pkg/polynomial"
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

	// pkg
	_ = container.Provide(polynomial.NewPolynomial)

	return container
}
