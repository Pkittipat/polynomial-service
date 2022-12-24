package handlers

import (
	"errors"
	"net/http"

	"github.com/Pkittipat/polynomial-service/internal/http/requests"
	"github.com/Pkittipat/polynomial-service/internal/http/responses"
	"github.com/Pkittipat/polynomial-service/internal/usecases"
	"github.com/gin-gonic/gin"
)

type PolynomialHandler interface {
	Calculate(c *gin.Context)
}

type polynomialHandler struct {
	usecase usecases.PolynomialUsecase
}

func NewPolynomialHandler(
	usecase usecases.PolynomialUsecase,
) PolynomialHandler {
	return &polynomialHandler{
		usecase: usecase,
	}
}

func (p *polynomialHandler) Calculate(c *gin.Context) {
	var request requests.CalculateForm
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isPolynimial := true
	if err := p.usecase.Calculate(request.X, request.Y, request.Z); err != nil {
		if !errors.Is(err, responses.ErrInvalidPolynomial) {
			responses.NewErrorResponse(err).Response(c, http.StatusInternalServerError)
		}
		responses.NewResponse(responses.NewCalculate(!isPolynimial)).
			Response(c, http.StatusOK)
		return
	}

	responses.NewResponse(responses.NewCalculate(isPolynimial)).
		Response(c, http.StatusOK)
}
