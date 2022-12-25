package usecases

import (
	"github.com/Pkittipat/polynomial-service/internal/http/responses"
	"github.com/Pkittipat/polynomial-service/pkg/polynomial"
)

type PolynomialUsecase interface {
	Guess(x, y, z int) error
	RandomArg() (datasetRes []int, degree, diff int)
}

type polynomialUsecase struct {
	pkg polynomial.Polynomial
}

func NewPolynomialUsecase(
	pkg polynomial.Polynomial,
) PolynomialUsecase {
	return &polynomialUsecase{
		pkg: pkg,
	}
}

func (p *polynomialUsecase) Guess(x, y, z int) error {
	_, isPolynomial := p.pkg.ProveXYZ(x, y, z)
	if !isPolynomial {
		return responses.ErrInvalidPolynomial
	}
	return nil
}

func (p *polynomialUsecase) RandomArg() (datasetRes []int, degree, diff int) {
	dataset := p.pkg.RandomArg()
	return dataset, p.pkg.GetDegree(), p.pkg.GetDiff()
}
