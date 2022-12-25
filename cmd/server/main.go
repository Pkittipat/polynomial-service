package main

import (
	"net/http"
	"strconv"

	"github.com/Pkittipat/polynomial-service/cmd"
	"github.com/Pkittipat/polynomial-service/config"
	"github.com/Pkittipat/polynomial-service/internal/http/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func main() {
	err := Run()
	if err != nil {
		panic(err)
	}
}

func Run() error {
	container := cmd.BuildContainer()

	var configInstance *config.Config
	if err := container.Invoke(func(_config *config.Config) {
		configInstance = _config
	}); err != nil {
		panic(err)
	}

	if configInstance.App.APP.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.New()
	setupMiddlewares(app, container)
	setupRoutes(app, container)

	err := app.Run(":" + strconv.FormatUint(uint64(configInstance.App.APP.Port), 10))
	return err
}

func setupMiddlewares(app *gin.Engine, container *dig.Container) {
}

func setupRoutes(app *gin.Engine, container *dig.Container) {
	var (
		polynomialHandler handlers.PolynomialHandler
	)

	if err := container.Invoke(func(
		_polynomialHandler handlers.PolynomialHandler,
	) {
		polynomialHandler = _polynomialHandler
	}); err != nil {
		panic(err)
	}

	app.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})

	polynomialGroup := app.Group("/polynomial")
	{
		polynomialGroup.POST("/calculate", polynomialHandler.Calculate)
		polynomialGroup.GET("/dataset", polynomialHandler.Dataset)
	}
}
